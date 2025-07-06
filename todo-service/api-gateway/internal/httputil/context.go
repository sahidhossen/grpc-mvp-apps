package httputil

import (
	"context"
	"net/http"
	"time"
)

const DefaultTimeout = 5 * time.Second

func WithTimeout(r *http.Request) (context.Context, context.CancelFunc) {
	return context.WithTimeout(r.Context(), DefaultTimeout)
}
