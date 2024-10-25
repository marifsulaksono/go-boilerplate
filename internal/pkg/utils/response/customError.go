package response

import (
	"fmt"
	"log"
)

type CustomError struct {
	StatusCode int
	Message    string
	Err        error
}

func (e *CustomError) Error() string {
	if e.Err != nil {
		log.Printf("StatusCode: %d, Message: %s, Detail: %v", e.StatusCode, e.Message, e.Err)
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
