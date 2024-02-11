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

func GetUserRoles(c echo.Context) ([]string, error) {
	roles := c.Get("roles")

	switch res := roles.(type) {
	case []string:
		return res, nil
	case []any:
		rolesStrSl := make([]string, 0, len(res))
		for _, v := range res {
			var role string
			var ok bool
			if role, ok = v.(string); !ok {
				return nil, errors.New("role must be string")
			}
			rolesStrSl = append(rolesStrSl, role)
		}
		return rolesStrSl, nil
	default:
		return nil, errors.New("roles is wrong type")
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
