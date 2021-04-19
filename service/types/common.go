package types

type StatusAndBodyFromResponse struct {
	// StatusAndBodyFromResponse Status Code
	StatusCode int32 `src:"status"`

	// StatusAndBodyFromResponse Status Message
	Status string `src:"status"`

	// StatusAndBodyFromResponse Body Content
	RawBody string `src:"body"`
}

//A response object that contains details about the error.
type Error struct {
	// Code of error.
	Code *int32 `json:"code,omitempty"`

	//Message of error.
	Message *string `json:"message,omitempty"`

	// Target
	Target *string `json:"target,omitempty"`
}
