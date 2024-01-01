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
