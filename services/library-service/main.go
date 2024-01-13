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
	container.MapRoute(di.PUT, "/authors/:id", routes.UpdateAuthor)
	container.MapRoute(di.DELETE, "/authors/:id", routes.DeleteAuthor)

	container.RunServer(config.Config.ServerOptions)
}
