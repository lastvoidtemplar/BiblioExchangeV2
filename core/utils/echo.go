package utils

import (
	"errors"

	"github.com/labstack/echo/v4"
)

func GetUserId(c echo.Context) (string, bool, error) {
	user_id := c.Get("userId")

	switch res := user_id.(type) {
	case string:
		return res, false, nil
	case nil:
		return "", true, nil
	default:
		return "", true, errors.New("userId is wrong type")
	}
}

type errorBody struct {
	Message string `json:"message"`
}

func ErrorHandler(c echo.Context, code int, err any) error {
	switch res := err.(type) {
	case error:
		return c.JSON(code, errorBody{
			Message: res.Error(),
		})
	case string:
		return c.JSON(code, errorBody{
			Message: res,
		})
	default:
		return c.String(code, "Invalid error for error handler")
	}

}
