package helper

import (
	"github.com/labstack/echo/v4"
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
