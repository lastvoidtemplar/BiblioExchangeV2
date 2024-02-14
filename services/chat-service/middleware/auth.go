package middleware_echo

import (
	"log"

	"github.com/labstack/echo/v4"
	"github.com/lastvoidtemplar/BiblioExchangeV2/core/authentication"
	authoptions "github.com/lastvoidtemplar/BiblioExchangeV2/core/authentication/auth_options"
	"github.com/lastvoidtemplar/BiblioExchangeV2/core/middleware"
)

func UseJWTHandlerMiddleware(c *echo.Echo, opt authoptions.AuthOptions) {
	rsaKey, err := authentication.LoadPublicKey(opt)

	if err != nil {
		log.Fatalf("Error: %e", err)
	}

	if err != nil {
		log.Fatalf("Error: rsaPublicKey - %s", err.Error())
	}

	c.Use(middleware.CreateJwtHandler(rsaKey))
}
