package grpchttpendpoint

import (
	"io"
	"net/http"

	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// DecodeJSONRequest decodes the JSON request into the given proto message reference.
// The HTTP error status code will be respond if decoding failed.
// UnmarshalOptions opts cannot be nil.
func DecodeJSONRequest(w http.ResponseWriter, r *http.Request, m proto.Message, opts *protojson.UnmarshalOptions) error {
	if nil == r.Body {
		http.Error(w, "empty request", http.StatusBadRequest)
		return ErrEmptyRequestBody
	}
	reqBody, err := io.ReadAll(r.Body)
	if nil != err {
		http.Error(w, "cannot load request", http.StatusBadRequest)
		return &ErrLoadRequest{
			Err: err,
		}
	}
	if err = opts.Unmarshal(reqBody, m); nil != err {
		http.Error(w, "malformed request", http.StatusBadRequest)
		return &ErrMalformedRequest{
			Err: err,
		}
	}
	return nil
}
