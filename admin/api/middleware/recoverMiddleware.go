package AdminMiddleware

import (
	"fmt"
	"net/http"
	"runtime/debug"
)

func RecoverMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				// Log the panic with stack trace
				stackTrace := string(debug.Stack())
				http.Error(w, fmt.Sprintf("Panic: %v\n\nStack trace:\n%s", err, stackTrace), http.StatusInternalServerError)

			}
		}()
		next.ServeHTTP(w, r)
	})
}
