package routes

import (
	"context"
	"database/sql"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	dbmodels "github.com/lastvoidtemplar/BiblioExchangeV2/core/db/models"
	"github.com/lastvoidtemplar/BiblioExchangeV2/core/db/queries"
	"github.com/lastvoidtemplar/BiblioExchangeV2/core/di"
	"github.com/lastvoidtemplar/BiblioExchangeV2/core/di/identificators"
)

func GetAuthorsPaginated(c *di.Container) echo.HandlerFunc {
	db, err := di.GetService[*sql.DB](c, identificators.Database)
	if err != nil {
		log.Fatalln(err)
	}
	ctx := context.Background()
	return func(c echo.Context) error {
		authors, err := queries.GetAllAuthors(ctx, db, 1, 10)
		if err != nil {
			return c.String(http.StatusInternalServerError, err.Error())
		}
		if authors == nil {
			return c.String(http.StatusNotFound, "empty")
		}
		return c.JSON(http.StatusOK, authors)
	}
}

type SingleAuthorDTO struct {
	dbmodels.Author
	CountStars int
}

func GetAuthorById(c *di.Container) echo.HandlerFunc {
	db, err := di.GetService[*sql.DB](c, identificators.Database)
	if err != nil {
		log.Fatalln(err)
	}
	ctx := context.Background()
	return func(c echo.Context) error {
		id := c.Param("id")

		author, err := queries.GetAuthorById(ctx, db, id)
		if err != nil {
			return c.String(http.StatusInternalServerError, err.Error())
		}
		if author == nil {
			return c.String(http.StatusNotFound, "empty")
		}

		if author.R == nil {
			return c.String(http.StatusNotFound, "R nil")
		}

		dto := SingleAuthorDTO{
			Author:     *author,
			CountStars: len(author.R.Authorpageratings),
		}
		return c.JSON(http.StatusOK, dto)
	}
}
