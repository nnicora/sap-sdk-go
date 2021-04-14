package types

type StatusAndBodyFromResponse struct {
	// StatusAndBodyFromResponse Status Code
	StatusCode int32 `src:"status"`

	// StatusAndBodyFromResponse Status Message
	Status string `src:"status"`

	// StatusAndBodyFromResponse Body Content
	RawBody string `src:"body"`
}
