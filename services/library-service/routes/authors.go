package routes

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	dbmodels "github.com/lastvoidtemplar/BiblioExchangeV2/core/db/models"
	"github.com/lastvoidtemplar/BiblioExchangeV2/core/db/queries"
	"github.com/lastvoidtemplar/BiblioExchangeV2/core/di"
	"github.com/lastvoidtemplar/BiblioExchangeV2/core/di/identificators"
	"github.com/lastvoidtemplar/BiblioExchangeV2/core/utils"
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
			return c.JSON(http.StatusOK, []struct{}{})
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
			return c.String(http.StatusNotFound, "Author with this id isn`t found!")
		}

		if author.R == nil {
			return c.String(http.StatusNotFound, "Author page ratings are missing!")
		}

		userId, anonymous, err := utils.GetUserId(c)

		if err != nil {
			return c.String(http.StatusInternalServerError,
				fmt.Sprintf("Error: %s!", err.Error()))
		}

		err = queries.AddAuthorPageView(ctx, db, id, userId, anonymous)

		if err != nil {
			log.Printf("Error: %s\n", err.Error())
			return c.String(http.StatusInternalServerError,
				fmt.Sprintf("Error when adding page view for author with id %s!", id))
		}

		dto := SingleAuthorDTO{
			Author:     *author,
			CountStars: len(author.R.Authorpageratings),
		}
		return c.JSON(http.StatusOK, dto)
	}
}
