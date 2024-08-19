package grpchttpendpoint

import (
	"net/http"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// ErrorStatusResponse generate HTTP error response based on given gRPC error status.
func ErrorStatusResponse(w http.ResponseWriter, err error) {
	w.Header().Set("Cache-Control", "no-cache, private")
	w.Header().Set("Pragma", "no-cache")
	st := status.Convert(err)
	httpCode := http.StatusInternalServerError
	httpMsg := "unknown error" // must not contain colon (":") character
	statusMsg := st.Message()
	switch st.Code() {
	case codes.Canceled:
		httpMsg = "canceled"
		httpCode = http.StatusServiceUnavailable
	case codes.Unknown:
		httpCode = http.StatusInternalServerError
		httpMsg = "unknown error"
	case codes.InvalidArgument:
		httpCode = http.StatusBadRequest
		httpMsg = "invalid argument"
	case codes.DeadlineExceeded:
		httpCode = http.StatusInternalServerError
		httpMsg = "deadline exceeded"
	case codes.NotFound:
		httpCode = http.StatusNotFound
		httpMsg = "not found"
	case codes.AlreadyExists:
		httpCode = http.StatusConflict
		httpMsg = "already exists"
	case codes.PermissionDenied:
		httpCode = http.StatusForbidden
		httpMsg = "permission denied"
	case codes.ResourceExhausted:
		httpCode = http.StatusServiceUnavailable
		httpMsg = "resource exhausted"
	case codes.FailedPrecondition:
		httpCode = http.StatusInternalServerError
		httpMsg = "failed precondition"
	case codes.Aborted:
		httpCode = http.StatusInternalServerError
		httpMsg = "aborted"
	case codes.OutOfRange:
		httpCode = http.StatusBadRequest
		httpMsg = "out of range"
	case codes.Unimplemented:
		httpCode = http.StatusNotImplemented
		httpMsg = "unimplemented"
	case codes.Internal:
		httpCode = http.StatusInternalServerError
		httpMsg = "internal error"
	case codes.Unavailable:
		httpCode = http.StatusServiceUnavailable
		httpMsg = "unavailable"
	case codes.DataLoss:
		httpCode = http.StatusInternalServerError
		httpMsg = "data loss"
	case codes.Unauthenticated:
		httpCode = http.StatusUnauthorized
		httpMsg = "unauthenticated"
	}
	if statusMsg != "" {
		httpMsg = httpMsg + ": " + statusMsg
	}
	http.Error(w, httpMsg, httpCode)
}
