package saperr

import (
	"fmt"
	"strings"
)

type SapError struct {
	Code    string
	Message string
	Orig    error
}

func New(code, message string, origin error) error {
	return &SapError{
		Code:    code,
		Message: message,
		Orig:    origin,
	}
}

func NewA(code string, messages ...string) error {
	return New(code, strings.Join(messages, ";"), nil)
}

func NewB(code string, origin error, messages ...string) error {
	return New(code, strings.Join(messages, ";"), origin)
}

func (e *SapError) Error() string {
	if e.Orig != nil {
		return fmt.Sprintf("%s, %s; %v", e.Code, e.Message, e.Orig)
	}
	return fmt.Sprintf("%s, %s", e.Code, e.Message)
}
