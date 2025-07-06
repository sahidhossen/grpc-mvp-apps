package httputil

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type ErrorResponse struct {
	Message string `json:"message"`
}

// HandleError is a generic helper to send an HTTP JSON error response.
func HandleError(w http.ResponseWriter, r *http.Request, logger *slog.Logger, err error, clientMessage string, statusCode int) {
	logger.Error("API request error",
		"error", err,
		"client_message", clientMessage,
		"status_code", statusCode,
		"path", r.URL.Path,
		"method", r.Method,
	)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	if err := json.NewEncoder(w).Encode(ErrorResponse{Message: clientMessage}); err != nil {
		logger.Error("Failed to encode error response", "error", err, "path", r.URL.Path)
	}
}

// HandleSuccess is a generic helper to send an HTTP JSON success response.
func HandleSuccess(w http.ResponseWriter, r *http.Request, logger *slog.Logger, data interface{}, statusCode int) {
	logger.Info("API request successful",
		"status_code", statusCode,
		"path", r.URL.Path,
		"method", r.Method,
	)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	if err := json.NewEncoder(w).Encode(data); err != nil {
		logger.Error("Failed to encode success response", "error", err, "path", r.URL.Path)
		http.Error(w, "Internal Server Error: Failed to encode response", http.StatusInternalServerError)
	}
}

// HandleGrpcError maps gRPC status codes to appropriate HTTP status codes and calls HandleError.
func HandleGrpcError(w http.ResponseWriter, r *http.Request, logger *slog.Logger, grpcErr error, defaultClientMessage string) {
	st, ok := status.FromError(grpcErr)
	if !ok {
		HandleError(w, r, logger, grpcErr, defaultClientMessage, http.StatusInternalServerError)
		return
	}

	switch st.Code() {
	case codes.NotFound:
		HandleError(w, r, logger, grpcErr, st.Message(), http.StatusNotFound)
	case codes.InvalidArgument:
		HandleError(w, r, logger, grpcErr, st.Message(), http.StatusBadRequest)
	case codes.AlreadyExists:
		HandleError(w, r, logger, grpcErr, st.Message(), http.StatusConflict)
	case codes.PermissionDenied:
		HandleError(w, r, logger, grpcErr, st.Message(), http.StatusForbidden)
	case codes.Unauthenticated:
		HandleError(w, r, logger, grpcErr, st.Message(), http.StatusUnauthorized)
	case codes.Unavailable:
		HandleError(w, r, logger, grpcErr, st.Message(), http.StatusServiceUnavailable)
	case codes.DeadlineExceeded:
		HandleError(w, r, logger, grpcErr, st.Message(), http.StatusGatewayTimeout)
	case codes.Canceled:
		HandleError(w, r, logger, grpcErr, st.Message(), http.StatusRequestTimeout)
	default:
		HandleError(w, r, logger, grpcErr, defaultClientMessage, http.StatusInternalServerError)
	}
}
