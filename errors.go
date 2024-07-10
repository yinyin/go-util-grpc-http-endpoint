package grpchttpendpoint

import (
	"errors"
)

// ErrEmptyRequestBody indicate the request body is empty.
var ErrEmptyRequestBody = errors.New("empty request body")

// ErrDecodeRequest indicate an error occur on decoding the request.
var ErrDecodeRequest = errors.New("cannot decode request")

// ErrLoadRequest indicate having issue on loading the request.
type ErrLoadRequest struct {
	Err error
}

func (e ErrLoadRequest) Error() string {
	return "cannot load request: " + e.Err.Error()
}

func (e ErrLoadRequest) Unwrap() []error {
	return []error{e.Err, ErrDecodeRequest}
}

// ErrMalformedRequest indicate the request is malformed and cannot be decode.
type ErrMalformedRequest struct {
	Err error
}

func (e ErrMalformedRequest) Error() string {
	return "malformed request: " + e.Err.Error()
}

func (e ErrMalformedRequest) Unwrap() []error {
	return []error{e.Err, ErrDecodeRequest}
}
