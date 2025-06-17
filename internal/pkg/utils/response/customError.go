package response

import (
	"fmt"
)

type CustomError struct {
	StatusCode int
	Message    string
	Err        error
}

func (e *CustomError) Error() string {
	if e.Err != nil {
		return fmt.Sprintf("%v", e.Err)
	}
	return ""
}

func (e *CustomError) Unwrap() error {
	return e.Err
}

func NewCustomError(statusCode int, message string, err error) *CustomError {
	return &CustomError{
		StatusCode: statusCode,
		Message:    message,
		Err:        err,
	}
}
