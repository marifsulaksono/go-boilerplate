package response

import (
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/marifsulaksono/go-echo-boilerplate/internal/constants"
)

type JSONResponse struct {
	Code    string      `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func BuildSuccessResponse(c echo.Context, statusCode int, message string, data any) error {
	return c.JSON(statusCode, JSONResponse{
		Code:    constants.RESPONSE_SUCCESS,
		Message: message,
		Data:    data,
	})
}

func BuildErrorResponse(c echo.Context, err error) error {
	var customErr *CustomError
	var statusCode int
	var message string
	var errorDetails interface{}

	if errors.As(err, &customErr) {
		statusCode = customErr.StatusCode
		message = customErr.Message
		errorDetails = customErr.Err
	} else {
		statusCode = http.StatusInternalServerError
		message = "An unexpected error occurred"
		errorDetails = err.Error()
	}

	// map the status code to a response code
	var code string
	switch statusCode {
	case http.StatusBadRequest:
		code = constants.RESPONSE_BAD_REQUEST
	case http.StatusNotFound:
		code = constants.RESPONSE_NOT_FOUND
	case http.StatusUnauthorized:
		code = constants.RESPONSE_UNAUTHENTICATED
	case http.StatusForbidden:
		code = constants.RESPONSE_UNAUTHORIZED
	case http.StatusTooManyRequests:
		code = constants.RESPONSE_TOO_MANY_REQUESTS
	case http.StatusInternalServerError:
		code = constants.RESPONSE_INTERNAL_SERVER_ERROR
	default:
		code = constants.RESPONSE_INTERNAL_SERVER_ERROR
	}

	response := JSONResponse{
		Code:    code,
		Message: message,
		Data:    errorDetails,
	}

	return c.JSON(statusCode, response)
}
