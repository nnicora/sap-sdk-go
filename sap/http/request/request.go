package request

import (
	"bytes"
	"context"
	"github.com/nnicora/sap-sdk-go/internal/processors"
	"github.com/nnicora/sap-sdk-go/internal/saperr"
	"github.com/nnicora/sap-sdk-go/internal/sapio"
	"github.com/nnicora/sap-sdk-go/sap"
	"github.com/nnicora/sap-sdk-go/sap/metainfo"
	"io"
	"net/http"
	"net/url"
	"reflect"
	"strings"
	"time"
)

type Request struct {
	RuntimeConfig *sap.RuntimeConfig
	ServiceInfo   metainfo.ServiceInfo

	Processors *processors.Processors
	Operation  *Operation

	CreationTime time.Time
	LastSignedAt time.Time
	AttemptTime  time.Time

	HTTPResponse        *http.Response
	ResponseBody        []byte
	ResponseBodyHandler func(int, []byte) ([]byte, error)
	OutputData          interface{}

	HTTPRequest   *http.Request
	requestBody   io.ReadSeeker
	streamingBody io.ReadCloser
	bodyStart     int64
	InputData     interface{}

	Error error

	//RetryCount int
	Retryable bool
	//RetryDelay time.Duration

	DisableFollowRedirects bool

	context context.Context
	built   bool

	safeBody *offsetReader
}

type Operation struct {
	Name string
	Http HTTP
}
type HTTP struct {
	Method      HTTPMethod
	Path        string
	UsePathAsIs bool
}

func New(ctx context.Context, cfg *sap.RuntimeConfig, serviceInfo metainfo.ServiceInfo, processors *processors.Processors,
	operation *Operation, params interface{}, data interface{}) *Request {

	httpReq, _ := createHttpRequest(ctx, &serviceInfo, operation)
	return &Request{
		RuntimeConfig: cfg,
		ServiceInfo:   serviceInfo,
		Processors:    processors,

		CreationTime: time.Now(),
		Operation:    operation,
		HTTPRequest:  httpReq,
		InputData:    params,
		OutputData:   data,

		context: ctx,
	}
}
func createHttpRequest(ctx context.Context, serviceInfo *metainfo.ServiceInfo, operation *Operation) (*http.Request, error) {
	httpOp := operation.Http
	httpReq, _ := http.NewRequestWithContext(ctx, httpOp.Method.String(), "", nil)

	path := buildPath(
		strings.TrimSuffix(serviceInfo.Endpoint.Host, "/"),
		serviceInfo.ServiceID,
		serviceInfo.APIVersion,
		strings.TrimPrefix(httpOp.Path, "/"),
	)
	if httpOp.UsePathAsIs {
		path = buildPath(
			strings.TrimSuffix(serviceInfo.Endpoint.Host, "/"),
			strings.TrimPrefix(httpOp.Path, "/"),
		)
	}

	if u, err := url.Parse(path); err != nil {
		httpReq.URL = &url.URL{}
		httpReq.URL.Path = serviceInfo.Endpoint.Host
	} else {
		httpReq.URL = u
	}

	return httpReq, nil
}

func buildPath(tokens ...string) string {
	elems := make([]string, 0)
	for _, v := range tokens {
		if len(v) > 0 {
			elems = append(elems, v)
		}
	}
	return strings.Join(elems, "/")
}

func (r *Request) HasError() bool {
	return r.Error != nil
}

func (r *Request) Context() context.Context {
	if r.context != nil {
		return r.context
	}
	return context.Background()
}

func (r *Request) Build() error {
	if !r.built {
		r.Processors.Using(Validate).Exec(r)
		if r.Error != nil {
			return r.Error
		}
		r.Processors.Using(Build).Exec(r)
		if r.Error != nil {
			return r.Error
		}
		r.built = true
	}

	return r.Error
}

func (r *Request) Sign() error {
	r.BuildAndSanitize()
	if r.Error != nil {
		return r.Error
	}
	r.Processors.Using(Sign).Exec(r)
	return r.Error
}

func (r *Request) BuildAndSanitize() error {
	r.Build()
	if r.Error != nil {
		return r.Error
	}

	sanitizeHostForHeader(r.HTTPRequest)
	return r.Error
}

func (r *Request) Send() error {
	defer func() {
		if r.OutputDataFilled() {
			r.readFromHttpResponseTo(r.OutputData)
		}

		r.Processors.Using(Complete).Exec(r)
	}()

	if err := r.Error; err != nil {
		return err
	}

	if r.InputDataFilled() {
		r.writeToHttpRequestFrom(r.InputData)
		//r.buildBody()
	}

	for {
		r.Error = nil
		r.AttemptTime = time.Now()

		r.Processors.Using(Validate).Exec(r)
		if r.Error != nil {
			return r.Error
		}

		if err := r.sendRequest(); err == nil {
			if r.Error != nil {
				return r.Error
			}
			return nil
		}
		r.Processors.Using(Retry).Exec(r)
		r.Processors.Using(AfterRetry).Exec(r)

		if err := r.prepareRetry(); err != nil {
			r.Error = err
			return err
		}
	}
}

func (r *Request) prepareRetry() error {
	r.HTTPRequest = copyHTTPRequest(r.HTTPRequest, nil)
	if r.Error != nil {
		return r.Error
	}

	// Closing response body to ensure that no response body is leaked
	// between retry attempts.
	if r.HTTPResponse != nil && r.HTTPResponse.Body != nil {
		r.HTTPResponse.Body.Close()
		r.ResponseBody = nil
	}

	return nil
}

func (r *Request) sendRequest() (sendErr error) {
	defer r.Processors.Using(CompleteAttempt).Exec(r)

	r.Retryable = false
	r.Processors.Using(Send).Exec(r)
	if r.Error != nil {
		return r.Error
	}

	r.Processors.Using(UnmarshalMeta).Exec(r)
	r.Processors.Using(ValidateResponse).Exec(r)
	if r.Error != nil {
		r.Processors.Using(UnmarshalError).Exec(r)
		return r.Error
	}

	r.Processors.Using(Unmarshal).Exec(r)
	if r.Error != nil {
		return r.Error
	}

	return nil
}

func (r *Request) InputDataFilled() bool {
	return r.InputData != nil && reflect.ValueOf(r.InputData).Elem().IsValid()
}
func (r *Request) OutputDataFilled() bool {
	return r.OutputData != nil && reflect.ValueOf(r.OutputData).Elem().IsValid()
}

// SetBufferBody will set the request's body bytes that will be sent to
// the service API.
func (r *Request) SetBytesBody(buf []byte) {
	r.SetReaderBody(bytes.NewReader(buf))
}

// SetStringBody sets the body of the request to be backed by a string.
func (r *Request) SetStringBody(s string) {
	r.SetReaderBody(strings.NewReader(s))
}

// SetReaderBody will set the request's body reader.
func (r *Request) SetReaderBody(reader io.ReadSeeker) {
	r.requestBody = reader

	if sapio.IsReaderSeekable(reader) {
		var err error
		// Get the Bodies current offset so retries will start from the same
		// initial position.
		r.bodyStart, err = reader.Seek(0, io.SeekCurrent)
		if err != nil {
			r.Error = saperr.New(saperr.Serialization,
				"failed to determine start of request body", err)
			return
		}
	}
	r.ResetBody()
}

func (r *Request) getNextRequestBody() (body io.ReadCloser, err error) {
	if r.streamingBody != nil {
		return r.streamingBody, nil
	}

	if r.safeBody != nil {
		r.safeBody.Close()
	}

	r.safeBody, err = newOffsetReader(r.requestBody, r.bodyStart)
	if err != nil {
		return nil, saperr.New(saperr.Serialization,
			"failed to get next request body reader", err)
	}

	// Go 1.8 tightened and clarified the rules code needs to use when building
	// requests with the http package. Go 1.8 removed the automatic detection
	// of if the Request.requestBody was empty, or actually had bytes in it. The SDK
	// always sets the Request.requestBody even if it is empty and should not actually
	// be sent. This is incorrect.
	//
	// Go 1.8 did add a http.NoBody value that the SDK can use to tell the http
	// client that the request really should be sent without a body. The
	// Request.requestBody cannot be set to nil, which is preferable, because the
	// field is exported and could introduce nil pointer dereferences for users
	// of the SDK if they used that field.
	//
	// Related golang/go#18257
	l, err := sapio.SeekerLen(r.requestBody)
	if err != nil {
		return nil, saperr.New(saperr.Serialization,
			"failed to compute request body size", err)
	}

	if l == 0 {
		body = NoBody
	} else if l > 0 {
		body = r.safeBody
	} else {
		// Hack to prevent sending bodies for methods where the body
		// should be ignored by the server. Sending bodies on these
		// methods without an associated ContentLength will cause the
		// request to socket timeout because the server does not handle
		// Transfer-Encoding: chunked bodies for these methods.
		//
		// This would only happen if a aws.ReaderSeekerCloser was used with
		// a io.Reader that was not also an io.Seeker, or did not implement
		// Len() method.
		switch r.Operation.Http.Method {
		case GET, HEAD, DELETE:
			body = NoBody
		default:
			body = r.safeBody
		}
	}

	return body, nil
}
