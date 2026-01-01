package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/nitesh111sinha/car-management/driver"
	carHandler "github.com/nitesh111sinha/car-management/handler/car"
	engineHandler "github.com/nitesh111sinha/car-management/handler/engine"
	carService "github.com/nitesh111sinha/car-management/service/car"
	engineService "github.com/nitesh111sinha/car-management/service/engine"
	carStore "github.com/nitesh111sinha/car-management/store/car"
	engineStore "github.com/nitesh111sinha/car-management/store/engine"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
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
		log.Fatal("Failed to execute schema: ", err)
	}

	router := mux.NewRouter()
	router.HandleFunc("/cars", carHandler.GetCars).Methods("GET")
	router.HandleFunc("/cars/{id:[0-9a-fA-F-]{36}}", carHandler.GetCarById).Methods("GET")
	router.HandleFunc("/cars/brand/{brand}", carHandler.GetCarByBrand).Methods("GET")
	router.HandleFunc("/cars", carHandler.CreateCar).Methods("POST")
	router.HandleFunc("/cars/{id:[0-9a-fA-F-]{36}}", carHandler.UpdateCar).Methods("PUT")
	router.HandleFunc("/cars/{id:[0-9a-fA-F-]{36}}", carHandler.DeleteCar).Methods("DELETE")

	router.HandleFunc("/engines", engineHandler.GetEngines).Methods("GET")
	router.HandleFunc("/engines/{id}", engineHandler.GetEngineById).Methods("GET")
	router.HandleFunc("/engines", engineHandler.CreateEngine).Methods("POST")
	router.HandleFunc("/engines/{id}", engineHandler.UpdateEngine).Methods("PUT")
	router.HandleFunc("/engines/{id}", engineHandler.DeleteEngine).Methods("DELETE")

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	addr := ":" + port

	log.Println("Server started on port", port)
	err = http.ListenAndServe(addr, router)
	if err != nil {
		panic(err)
	}
}

func executeSchema(db *sql.DB, schemaFile string) any {
	sqlFile, err := os.ReadFile(schemaFile)
	if err != nil {
		return err
	}
	_, err = db.Exec(string(sqlFile))
	if err != nil {
		return err
	}
	return nil
}
