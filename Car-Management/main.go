package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"

	"github.com/nitesh111sinha/car-management/driver"
	carHandler "github.com/nitesh111sinha/car-management/handler/car"
	engineHandler "github.com/nitesh111sinha/car-management/handler/engine"
	"github.com/nitesh111sinha/car-management/handler/login"
	"github.com/nitesh111sinha/car-management/middleware"
	carService "github.com/nitesh111sinha/car-management/service/car"
	engineService "github.com/nitesh111sinha/car-management/service/engine"
	carStore "github.com/nitesh111sinha/car-management/store/car"
	engineStore "github.com/nitesh111sinha/car-management/store/engine"

	"go.opentelemetry.io/contrib/instrumentation/github.com/gorilla/mux/otelmux"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.24.0"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	if err := godotenv.Load(); err != nil {
		panic(err)
	}

	traceProvider, err := startTracing()
	if err != nil {
		panic(err)
	}
	defer func() {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		if err := traceProvider.Shutdown(ctx); err != nil {
			log.Println("Trace shutdown error:", err)
		}
	}()

	otel.SetTracerProvider(traceProvider)

	driver.InitDB()
	defer driver.CloseDB()

	db := driver.GetDB()

	carStore := carStore.NewCarStore(db)
	engineStore := engineStore.NewEngineStore(db)

	carService := carService.NewCarService(carStore)
	engineService := engineService.NewEngineService(engineStore)

	carHandler := carHandler.NewCarHandler(carService)
	engineHandler := engineHandler.NewEngineHandler(engineService)

	schemaFile := os.Getenv("SCHEMA_FILE")
	if schemaFile == "" {
		panic("SCHEMA_FILE not set")
	}

	if err := executeSchema(db, schemaFile); err != nil {
		log.Fatal("Failed to execute schema:", err)
	}

	router := mux.NewRouter()
	router.Use(otelmux.Middleware("car-management"))
	router.Use(middleware.MetricsMiddleware)

	router.HandleFunc("/login", login.LoginHandler).Methods("POST")

	protected := router.PathPrefix("/").Subrouter()
	protected.Use(middleware.AuthMiddleware)

	protected.HandleFunc("/cars", carHandler.GetCars).Methods("GET")
	protected.HandleFunc("/cars/{id:[0-9a-fA-F-]{36}}", carHandler.GetCarById).Methods("GET")
	protected.HandleFunc("/cars/brand/{brand}", carHandler.GetCarByBrand).Methods("GET")
	protected.HandleFunc("/cars", carHandler.CreateCar).Methods("POST")
	protected.HandleFunc("/cars/{id:[0-9a-fA-F-]{36}}", carHandler.UpdateCar).Methods("PUT")
	protected.HandleFunc("/cars/{id:[0-9a-fA-F-]{36}}", carHandler.DeleteCar).Methods("DELETE")

	protected.HandleFunc("/engines", engineHandler.GetEngines).Methods("GET")
	protected.HandleFunc("/engines/{id}", engineHandler.GetEngineById).Methods("GET")
	protected.HandleFunc("/engines", engineHandler.CreateEngine).Methods("POST")
	protected.HandleFunc("/engines/{id}", engineHandler.UpdateEngine).Methods("PUT")
	protected.HandleFunc("/engines/{id}", engineHandler.DeleteEngine).Methods("DELETE")

	router.Handle("/metrics", promhttp.Handler())

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Println("Server started on port", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}

func executeSchema(db *sql.DB, schemaFile string) error {
	sqlFile, err := os.ReadFile(schemaFile)
	if err != nil {
		return err
	}
	_, err = db.Exec(string(sqlFile))
	return err
}

func startTracing() (*sdktrace.TracerProvider, error) {
	exporter, err := otlptracehttp.New(
		context.Background(),
		otlptracehttp.WithEndpoint("jaeger:4318"),
		otlptracehttp.WithInsecure(),
	)
	if err != nil {
		return nil, fmt.Errorf("creating OTLP exporter: %w", err)
	}

	res, err := resource.New(
		context.Background(),
		resource.WithSchemaURL(semconv.SchemaURL),
		resource.WithAttributes(
			semconv.ServiceNameKey.String("car-management"),
		),
	)
	if err != nil {
		return nil, err
	}

	tp := sdktrace.NewTracerProvider(
		sdktrace.WithBatcher(exporter),
		sdktrace.WithResource(res),
		sdktrace.WithSampler(sdktrace.AlwaysSample()),
	)

	return tp, nil
}
