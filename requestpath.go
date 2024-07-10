package grpchttpendpoint

import (
	"net/http"
)

// RequestPathInBytes returns the URL path of request in bytes.
// Leading slashes will be removed.
func RequestPathInBytes(r *http.Request) []byte {
	result := []byte(r.URL.Path)
	for (len(result) > 0) && (result[0] == '/') {
		result = result[1:]
	}
	return result
}
