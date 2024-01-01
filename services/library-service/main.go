package main

import (
	"github.com/lastvoidtemplar/BiblioExchangeV2/core/di"
	"github.com/lastvoidtemplar/BiblioExchangeV2/library-service/config"
	"github.com/lastvoidtemplar/BiblioExchangeV2/library-service/routes"
)

func main() {
	container := di.New().
		AddDatabase(config.Config.DatabaseOptions).
		Build()

	container.UseJWTHandlerMiddleware(config.Config.AuthOptions)

	container.MapRoute(di.GET, "/authors", routes.GetAuthorsPaginated)
	container.MapRoute(di.GET, "/authors/:id", routes.GetAuthorById)
	container.MapRoute(di.POST, "/authors", routes.CreateAuthor)

	container.RunServer(config.Config.ServerOptions)
}
