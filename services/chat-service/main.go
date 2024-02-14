package main

import (
	"chat-service/config"
	"chat-service/routes"

	"github.com/lastvoidtemplar/BiblioExchangeV2/core/di"
)

type Message struct {
	UserID  string `json:"userId"`
	Content string `json:"content"`
}

func main() {
	container := di.New().
		AddDatabase(config.Config.DatabaseOptions).
		Build()

	container.UseJWTHandlerMiddleware(config.Config.AuthOptions)

	container.MapRoute(di.GET, "/ws", routes.HandleWebSocket)
	container.MapRoute(di.GET, "/messages/:id", routes.GetMessagesMessage)
	container.MapRoute(di.POST, "/messages", routes.PostMessage)

	container.RunServer(config.Config.ServerOptions)
}
