package grpchttpendpoint

import (
	"net/http"
	"strconv"

	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

func attachJSONContentHeader(w http.ResponseWriter, b []byte) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("Content-Length", strconv.FormatInt(int64(len(b)), 10))
}

// JSONResponseWithStatusCode generate JSON response based on given protobuf
// message m with given statusCode and private Cache-Control headers enabled.
// MarshalOptions opts cannot be nil.
func JSONResponseWithStatusCode(w http.ResponseWriter, m proto.Message, statusCode int, opts *protojson.MarshalOptions) (err error) {
	b, err := opts.Marshal(m)
	if err != nil {
		http.Error(w, "500 JSONResponseWithStatusCode Failed:\n"+err.Error(), http.StatusInternalServerError)
		return err
	}
	attachJSONContentHeader(w, b)
	w.Header().Set("Cache-Control", "no-cache, private")
	w.Header().Set("Pragma", "no-cache")
	w.WriteHeader(statusCode)
	w.Write(b)
	return nil
}
