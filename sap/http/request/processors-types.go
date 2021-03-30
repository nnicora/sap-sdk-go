package request

import "github.com/nnicora/sap-sdk-go/internal/processors"

const (
	Validate processors.Type = iota
	Build
	BuildStream
	Sign
	Send
	ValidateResponse
	Unmarshal
	UnmarshalStream
	UnmarshalMeta
	UnmarshalError
	Retry
	AfterRetry
	CompleteAttempt
	Complete
)
