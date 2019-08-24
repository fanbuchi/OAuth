package prometheus

import (
	"net/http"

	prometheus "github.com/prometheus/client_golang/prometheus/promhttp"
)

// MetricsWrapper prometheus
func MetricsWrapper(h http.Handler) http.Handler {
	ph := prometheus.Handler()
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/metrics" {
			ph.ServeHTTP(w, r)
			return
		}

		h.ServeHTTP(w, r)
	})
}
