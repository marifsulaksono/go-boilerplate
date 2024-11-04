package helper

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/marifsulaksono/go-echo-boilerplate/internal/pkg/utils/response"
)

const E_INVALID_FORMAT = "Invalid Field Format"
const E_REQUIRED = "Invalid Required Field"

func BindRequest(c echo.Context, payload interface{}, skipValidation bool) error {
	if err := c.Bind(payload); err != nil {
		return echo.NewHTTPError(400, E_INVALID_FORMAT)
	}

	if !skipValidation {
		if err := c.Validate(payload); err != nil {
			return echo.NewHTTPError(422, err.Error())
		}
	}

	return nil
}

func GetPayloadAndRecycle(c echo.Context) (interface{}, error) {
	if c.Request().Method == http.MethodDelete {
		pathSegments := strings.Split(c.Request().URL.Path, "/")
		return pathSegments[len(pathSegments)-1], nil
	}

	// this method is the same as c.Bind
	bodyBytes, err := io.ReadAll(c.Request().Body)
	if err != nil {
		return nil, response.NewCustomError(http.StatusBadRequest, "Failed to read request body", nil)
	}

	// refill the body so it can be read again
	c.Request().Body = io.NopCloser(bytes.NewBuffer(bodyBytes))

	var payload map[string]interface{}
	if err := json.Unmarshal(bodyBytes, &payload); err != nil {
		return nil, response.NewCustomError(http.StatusBadRequest, "Failed to parse payload", nil)
	}
	return payload, nil
}
