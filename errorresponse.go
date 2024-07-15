package grpchttpendpoint

import (
	"net/http"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// ErrorStatusResponse generate HTTP error response based on given error.
func ErrorStatusResponse(w http.ResponseWriter, err error) {
	w.Header().Set("Cache-Control", "no-cache, private")
	w.Header().Set("Pragma", "no-cache")
	switch status.Code(err) {
	case codes.Canceled:
		http.Error(w, "canceled", http.StatusServiceUnavailable)
	case codes.Unknown:
		http.Error(w, "unknown error", http.StatusInternalServerError)
	case codes.InvalidArgument:
		http.Error(w, "invalid argument", http.StatusBadRequest)
	case codes.DeadlineExceeded:
		http.Error(w, "deadline exceeded", http.StatusInternalServerError)
	case codes.NotFound:
		http.Error(w, "not found", http.StatusNotFound)
	case codes.AlreadyExists:
		http.Error(w, "already exists", http.StatusConflict)
	case codes.PermissionDenied:
		http.Error(w, "permission denied", http.StatusForbidden)
	case codes.ResourceExhausted:
		http.Error(w, "resource exhausted", http.StatusServiceUnavailable)
	case codes.FailedPrecondition:
		http.Error(w, "failed precondition", http.StatusInternalServerError)
	case codes.Aborted:
		http.Error(w, "aborted", http.StatusInternalServerError)
	case codes.OutOfRange:
		http.Error(w, "out of range", http.StatusBadRequest)
	case codes.Unimplemented:
		http.Error(w, "unimplemented", http.StatusNotImplemented)
	case codes.Internal:
		http.Error(w, "internal error", http.StatusInternalServerError)
	case codes.Unavailable:
		http.Error(w, "unavailable", http.StatusServiceUnavailable)
	case codes.DataLoss:
		http.Error(w, "data loss", http.StatusInternalServerError)
	case codes.Unauthenticated:
		http.Error(w, "unauthenticated", http.StatusUnauthorized)
	default:
		http.Error(w, "unknown error", http.StatusInternalServerError)
	}
}
