package utils

import (
	"errors"
	"log"

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
	error string
}

func ErrorHandler(c echo.Context, code int, err any) error {
	switch res := err.(type) {
	case error:
		log.Panicf("Error: %s", res.Error())
		return c.JSON(code, errorBody{
			error: res.Error(),
		})
	case string:
		return c.JSON(code, errorBody{
			error: res,
		})
	default:
		return c.String(code, "Invalid error for error handler")
	}

}
