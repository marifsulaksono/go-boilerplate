package helper

import (
	"github.com/labstack/echo/v4"
)

const E_INVALID_FORMAT = "Invalid Field Format"
const E_REQUIRED = "Invalid Required Field"

func BindRequest(c echo.Context, payload interface{}, skipValidation bool) error {
	if err := c.Bind(payload); err != nil {
		return err
	}

	if !skipValidation {
		if err := c.Validate(payload); err != nil {
			return err
		}
	}

	return nil
}
