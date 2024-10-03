package AdminMiddleware

import (
	"net/http"

	"github.com/go-chi/chi/middleware"
)

// CustomLogger skips logging for specific endpoints like /health
func CustomLogger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Skip logging for the health check endpoint
		if r.URL.Path == "/health" {
			next.ServeHTTP(w, r)
			return
		}

		// Log all other requests using the default logger
		middleware.Logger(next).ServeHTTP(w, r)
	})
}
