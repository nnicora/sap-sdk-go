// +build go1.8

package request

import (
	"github.com/nnicora/sap-sdk-go/internal/saperr"
	"net/http"
)

// NoBody is a http.NoBody reader instructing Go HTTP client to not include
// and body in the HTTP request.
var NoBody = http.NoBody

func (r *Request) ResetBody() {
	body, err := r.getNextRequestBody()
	if err != nil {
		r.Error = saperr.New(saperr.Serialization,
			"failed to reset request body", err)
		return
	}

	r.HTTPRequest.Body = body
	r.HTTPRequest.GetBody = r.getNextRequestBody
}
