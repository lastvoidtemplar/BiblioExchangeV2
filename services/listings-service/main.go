package main

import (
	"listings-service/config"
	"listings-service/routes"
	"log"

	"github.com/lastvoidtemplar/BiblioExchangeV2/core/di"
	"github.com/lastvoidtemplar/BiblioExchangeV2/core/di/identificators"
	"github.com/lastvoidtemplar/BiblioExchangeV2/core/upload"
)

func main() {

	container := di.New().
		AddDatabase(config.Config.DatabaseOptions).
		AddUploadService().
		Build()

	container.UseJWTHandlerMiddleware(config.Config.AuthOptions)
	container.MapRoute(di.POST, "/listings", routes.CreateListing)
	container.MapRoute(di.GET, "/listings/:id", routes.GetBooksPaginated)
	container.InitService(func(c *di.Container) {
		uploadService, err := di.GetService[*upload.UploadService](c, identificators.UploadService)
		if err != nil {
			log.Fatalln(err)
		}
		uploadService.Init(config.Config.UploadOptions)
	})
	container.RunServer(config.Config.ServerOptions)
}
