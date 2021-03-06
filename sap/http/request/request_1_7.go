// +build !go1.8

package request

import "io"

// NoBody is an io.ReadCloser with no bytes. Read always returns EOF
// and Close always returns nil. It can be used in an outgoing client
// request to explicitly signal that a request has zero bytes.
// An alternative, however, is to simply set Request.requestBody to nil.
//
// Copy of Go 1.8 NoBody type from net/http/http.go
type noBody struct{}

func (noBody) Read([]byte) (int, error)         { return 0, io.EOF }
func (noBody) Close() error                     { return nil }
func (noBody) WriteTo(io.Writer) (int64, error) { return 0, nil }

// NoBody is an empty reader that will trigger the Go HTTP client to not include
// and body in the HTTP request.
var NoBody = noBody{}

func (r *Request) ResetBody() {
	body, err := r.getNextRequestBody()
	if err != nil {
		r.Error = err
		return
	}

	r.HTTPRequest.Body = body
}
