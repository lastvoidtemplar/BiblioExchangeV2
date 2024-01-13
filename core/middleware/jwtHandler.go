package middleware

import (
	"crypto/rsa"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

func CreateJwtHandler(rsaPublicKey *rsa.PublicKey) func(next echo.HandlerFunc) echo.HandlerFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			if c.Request().Header["Authorization"] != nil {
				tokenStr := strings.Split(c.Request().Header["Authorization"][0], " ")
				if len(tokenStr) == 2 && tokenStr[0] == "Bearer" {
					claims := jwt.MapClaims{}
					token, err := jwt.ParseWithClaims(tokenStr[1], &claims, func(token *jwt.Token) (interface{}, error) {
						return rsaPublicKey, nil
					})

					if err != nil {
						c.String(http.StatusBadRequest, err.Error())
						return nil
					}

					if !token.Valid {
						c.String(http.StatusBadRequest, "invalid token")
						return nil
					}

					c.Set("userId", claims["sub"])
				}

			}
			return next(c)
		}
	}
}
