package middleware

import (
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
)

var (
	requestCounter = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_requests_total",
			Help: "Total number of HTTP requests",
		},
		[]string{"path", "method", "status_code"},
	)

	requestDuration = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "http_request_duration_seconds",
			Help:    "Duration of HTTP requests",
			Buckets: []float64{0.1, 0.3, 0.5, 0.7, 0.9, 1.0},
		},
		[]string{"path", "method"},
	)

	statusCounter = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_response_status_total",
			Help: "Total number of HTTP responses by status code",
		},
		[]string{"path", "method", "status_code"},
	)
)

type responseWriter struct {
	http.ResponseWriter
	statusCode int
}

func (rw *responseWriter) WriteHeader(statusCode int) {
	rw.statusCode = statusCode
	rw.ResponseWriter.WriteHeader(statusCode)
}

func init() {
	prometheus.MustRegister(requestCounter)
	prometheus.MustRegister(requestDuration)
	prometheus.MustRegister(statusCounter)
}

func MetricsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// start time of the request
		start := time.Now()

		ww := &responseWriter{w, http.StatusOK}

		next.ServeHTTP(ww, r)

		// 	duration of the request
		duration := time.Since(start)

		// increment the request counter
		requestCounter.WithLabelValues(r.URL.Path, r.Method, string(ww.statusCode)).Inc()

		// increment the status counter
		statusCounter.WithLabelValues(r.URL.Path, r.Method, string(ww.statusCode)).Inc()

		// record the request duration
		requestDuration.WithLabelValues(r.URL.Path, r.Method).Observe(duration.Seconds())
	})
}
