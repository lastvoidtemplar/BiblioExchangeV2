package main

import (
	"database/sql"
	"log"
	"upload-service/config"
	"upload-service/routes"
	"upload-service/upload"

	"github.com/lastvoidtemplar/BiblioExchangeV2/core/di"
	"github.com/lastvoidtemplar/BiblioExchangeV2/core/di/identificators"
)

func main() {
	container := di.New().
		AddDatabase(config.Config.DatabaseOptions).
		RegisterService(upload.UploadServerIdentificator, upload.New(config.Config.MinioOptions)).
		Build()

	container.UseJWTHandlerMiddleware(config.Config.AuthOptions)
	container.MapRoute(di.POST, "/upload/:id", routes.CreateListing)
	container.InitService(func(c *di.Container) {
		db, err := di.GetService[*sql.DB](c, identificators.Database)
		if err != nil {
			log.Fatalln(err)
		}

		uploadServer, err := di.GetService[*upload.UploadServerServer](c,
			identificators.Identificator(upload.UploadServerIdentificator))
		if err != nil {
			log.Fatalln(err)
		}

		uploadServer.Init(db, config.Config.GrpcServerOptions)
	})
	container.RunServer(config.Config.ServerOptions)
}
