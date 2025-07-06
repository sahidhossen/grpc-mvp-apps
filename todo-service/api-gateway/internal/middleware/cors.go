package middleware

import (
	"log/slog"
	"net/http"
)

// corsMiddleware adds CORS headers to allow requests from the React frontend.
func CORSMiddleware(logger *slog.Logger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Set common CORS headers for all responses
			w.Header().Set("Access-Control-Allow-Origin", "*") // For development, "*" is fine. In prod, specify client origins.
			w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS, PATCH")
			w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Accept")
			w.Header().Set("Access-Control-Expose-Headers", "Content-Length")
			w.Header().Set("Access-Control-Max-Age", "86400")

			// Handle preflight OPTIONS requests
			if r.Method == http.MethodOptions {
				logger.Info("CORS: Handling preflight OPTIONS request", "path", r.URL.Path, "origin", r.Header.Get("Origin"))
				w.WriteHeader(http.StatusNoContent)
				return
			}
			next.ServeHTTP(w, r)
		})
	}
}
