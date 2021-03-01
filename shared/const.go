package shared

import "errors"

var (
	ErrNilAuthor = errors.New("Author not found")
	ErrNilBook   = errors.New("Book not found")
)

type DefaultResponse struct {
	Code    int    `json:"code"`
	Status  bool   `json:"status"`
	Message string `json:"message,omitempty"`
}

type ErrorResponse struct {
	DefaultResponse
	Error []string `json:"errors"`
}
