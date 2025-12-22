package metrics

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	httpRequestsTotal = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_requests_total",
			Help: "Total number of HTTP requests.",
		},
		[]string{"path", "method", "status"},
	)

	httpRequestDuration = promauto.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "http_request_duration_seconds",
			Help:    "Duration of HTTP requests in seconds.",
			Buckets: prometheus.DefBuckets,
		},
		[]string{"path", "method"},
	)

	handlerCalls = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Name: "handler_calls_total",
			Help: "Number of handler executions grouped by result.",
		},
		[]string{"handler", "result"},
	)
)

type statusRecorder struct {
	http.ResponseWriter
	status int
}

func (r *statusRecorder) WriteHeader(code int) {
	r.status = code
	r.ResponseWriter.WriteHeader(code)
}

// HTTPMetricsMiddleware collects count and latency per route.
func HTTPMetricsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		recorder := &statusRecorder{ResponseWriter: w, status: http.StatusOK}
		next.ServeHTTP(recorder, r)

		path := r.URL.Path
		if route := mux.CurrentRoute(r); route != nil {
			if tpl, err := route.GetPathTemplate(); err == nil {
				path = tpl
			}
		}

		httpRequestsTotal.WithLabelValues(path, r.Method, strconv.Itoa(recorder.status)).Inc()
		httpRequestDuration.WithLabelValues(path, r.Method).Observe(time.Since(start).Seconds())
	})
}

// Handler exposes Prometheus metrics endpoint.
func Handler() http.Handler {
	return promhttp.Handler()
}

// CountHandler increments per-handler success/error counters.
func CountHandler(handler, result string) {
	handlerCalls.WithLabelValues(handler, result).Inc()
}
