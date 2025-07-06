package middleware

import (
	"log/slog"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

// LoggingMiddleware is a simple example of a middleware function.
func LoggingMiddleware(logger *slog.Logger) mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()
			next.ServeHTTP(w, r)
			logger.Info("HTTP Request",
				"method", r.Method,
				"uri", r.RequestURI,
				"protocol", r.Proto,
				"duration", time.Since(start),
			)
		})
	}
}
