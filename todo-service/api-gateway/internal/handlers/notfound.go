package handlers

import (
	"net/http"
)

// NotFoundHandler provides a common fallback for unmatched routes.
// It handles OPTIONS preflights with 204 No Content and other methods with 404 Not Found.
// This assumes a global CORS middleware has already set necessary headers.
func NotFoundHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodOptions {
		// For OPTIONS preflight requests that fall through all other routes,
		// respond with 204 No Content. CORS middleware should have already
		// added the Access-Control-Allow-* headers.
		w.WriteHeader(http.StatusNoContent)
		return
	}
	// For any other HTTP method that didn't match a specific route,
	// respond with 404 Not Found.
	http.NotFound(w, r)
}
