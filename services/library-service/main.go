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

	container.MapRoute(di.POST, "/authors/:id/star", routes.ToggleAuthorStarRating)

	container.MapRoute(di.POST, "/authors/:id/review", routes.CreateAuthorReview)
	container.MapRoute(di.PUT, "/authors/reviews/:id", routes.UpdateAuthorReview)
	container.MapRoute(di.DELETE, "/authors/reviews/:id", routes.DeleteAuthorReview)

	container.MapRoute(di.GET, "/authors/:id/books", routes.GetBooksPaginated)
	container.MapRoute(di.POST, "/authors/:id/books", routes.CreateBookById)
	container.MapRoute(di.GET, "/books/:id", routes.GetBookById)
	container.MapRoute(di.PUT, "/books/:id", routes.UpdateBook)
	container.MapRoute(di.DELETE, "/books/:id", routes.DeleteBook)

	container.MapRoute(di.POST, "/books/:id/star", routes.ToggleBookStarRating)

	container.MapRoute(di.POST, "/books/:id/review", routes.CreateBookReview)
	container.MapRoute(di.PUT, "/books/reviews/:id", routes.UpdateBookReview)
	container.MapRoute(di.DELETE, "/books/reviews/:id", routes.DeleteBookReview)

	container.RunServer(config.Config.ServerOptions)
}
