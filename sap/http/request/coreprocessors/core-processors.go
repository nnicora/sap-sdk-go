package coreprocessors

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/nnicora/sap-sdk-go/internal/processors"
	"github.com/nnicora/sap-sdk-go/internal/saperr"
	"github.com/nnicora/sap-sdk-go/internal/sapio"
	"github.com/nnicora/sap-sdk-go/sap/http/request"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"regexp"
	"strconv"
)

// BuildContentLengthProcessor builds the content length of a request based on the body,
// or will use the HTTPRequest.Header's "Content-Length" if defined. If unable
// to determine request body length and no "Content-Length" was specified it will panic.
//
// The Content-Length will only be added to the request if the length of the body
// is greater than 0. If the body is empty or the current `Content-Length`
// header is <= 0, the header will also be stripped.
var BuildContentLengthProcessor = processors.DefaultProcessor{
	Name: "core.BuildContentLengthProcessor",
	Handler: func(t interface{}) {
		r := t.(*request.Request)
		var length int64

		if slength := r.HTTPRequest.Header.Get("Content-Length"); slength != "" {
			length, _ = strconv.ParseInt(slength, 10, 64)
		} else {
			if r.HTTPRequest.Body != nil {
				var err error
				if sapio.IsReaderSeekable(r.HTTPRequest.Body) {
					length, err = sapio.SeekerLen(r.HTTPRequest.Body.(io.ReadSeeker))
				} else {
					err = errors.New("request body is not seekable")
				}
				if err != nil {
					r.Error = saperr.New(saperr.Serialization, "failed to get request body's length", err)
					return
				}
			}
		}

		if length > 0 {
			r.HTTPRequest.ContentLength = length
			r.HTTPRequest.Header.Set("Content-Length", fmt.Sprintf("%d", length))
		} else {
			r.HTTPRequest.ContentLength = 0
			r.HTTPRequest.Header.Del("Content-Length")
		}
	},
}

var reStatusCode = regexp.MustCompile(`^(\d{3})`)

// ValidateReqSigProcessor is a request handler to request's are valid
var ValidateReqSigProcessor = processors.DefaultProcessor{
	Name: "core.ValidateReqSigProcessor",
	Handler: func(t interface{}) {
		r, _ := t.(*request.Request)
		r.Sign()
	},
}

// SendProcessor is a request handler to send service request using HTTP client.
var SendProcessor = processors.DefaultProcessor{
	Name: "core.SendProcessor",
	Handler: func(t interface{}) {
		r := t.(*request.Request)
		sender := sendFollowRedirects
		if r.DisableFollowRedirects {
			sender = sendWithoutFollowRedirects
		}

		if request.NoBody == r.HTTPRequest.Body {
			// Strip off the request body if the NoBody reader was used as a
			// place holder for a request body. This prevents the SDK from
			// making requests with a request body when it would be invalid
			// to do so.
			//
			// Using a shallow copy of the http.Request to ensure the race condition
			// of transport on requestBody will not trigger
			reqOrig, reqCopy := r.HTTPRequest, *r.HTTPRequest
			reqCopy.Body = nil
			r.HTTPRequest = &reqCopy
			defer func() {
				r.HTTPRequest = reqOrig
			}()
		}

		var err error
		r.HTTPResponse, err = sender(r)
		if err != nil {
			handleSendError(r, err)
		} else {
			r.ResponseBody, _ = ioutil.ReadAll(r.HTTPResponse.Body)
		}
	},
}

func sendFollowRedirects(r *request.Request) (*http.Response, error) {
	if httClient, err := r.ServiceInfo.EndpointHttpClient(); err != nil {
		return nil, err
	} else {
		return httClient.Do(r.HTTPRequest)
	}
}

func sendWithoutFollowRedirects(r *request.Request) (*http.Response, error) {
	if httClient, err := r.ServiceInfo.EndpointHttpClient(); err != nil {
		return nil, err
	} else {
		transport := httClient.Transport
		if transport == nil {
			transport = http.DefaultTransport
		}
		return transport.RoundTrip(r.HTTPRequest)
	}
}

func handleSendError(r *request.Request, err error) {
	// Prevent leaking if an HTTPResponse was returned. Clean up
	// the body.
	if r.HTTPResponse != nil {
		r.HTTPResponse.Body.Close()
		r.ResponseBody = nil
	}
	// Capture the case where url.Error is returned for error processing
	// response. e.g. 301 without location header comes back as string
	// error and r.HTTPResponse is nil. Other URL redirect errors will
	// comeback in a similar method.
	if e, ok := err.(*url.Error); ok && e.Err != nil {
		if s := reStatusCode.FindStringSubmatch(e.Err.Error()); s != nil {
			code, _ := strconv.ParseInt(s[1], 10, 64)
			r.HTTPResponse = &http.Response{
				StatusCode: int(code),
				Status:     http.StatusText(int(code)),
				Body:       ioutil.NopCloser(bytes.NewReader([]byte{})),
			}
			return
		}
	}
	if r.HTTPResponse == nil {
		// Add a dummy request response object to ensure the HTTPResponse
		// value is consistent.
		r.HTTPResponse = &http.Response{
			StatusCode: 0,
			Status:     http.StatusText(0),
			Body:       ioutil.NopCloser(bytes.NewReader([]byte{})),
		}
	}
	// Catch all request errors, and let the default retrier determine
	// if the error is retryable.
	r.Error = saperr.New("RequestError", "send request failed", err)

	// Override the error with a context canceled error, if that was canceled.
	ctx := r.Context()
	select {
	case <-ctx.Done():
		r.Error = fmt.Errorf("CanceledErrorCode, request context canceled; %s", ctx.Err())
		r.Retryable = false
	default:
	}
}

var ValidateResponseProcessor = processors.DefaultProcessor{
	Name: "core.ValidateResponseProcessor",
	Handler: func(t interface{}) {
		r := t.(*request.Request)
		if r.HTTPResponse.StatusCode == 0 || r.HTTPResponse.StatusCode >= 300 {
			// this may be replaced by an UnmarshalError handler
			r.Error = saperr.NewA(r.HTTPResponse.Status, r.Operation.Name)
		}
	},
}

var ValidateEndpointProcessor = processors.DefaultProcessor{
	Name: "core.ValidateEndpointProcessor",
	Handler: func(t interface{}) {
		r := t.(*request.Request)
		if r.ServiceInfo.Endpoint.Host == "" {
			r.Error = saperr.NewA("MissingEndpoint", "'Endpoint' configuration is required for this service")
		}
	},
}
