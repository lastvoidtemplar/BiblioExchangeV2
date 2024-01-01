package routes

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/lastvoidtemplar/BiblioExchangeV2/core/db/queries"
	"github.com/lastvoidtemplar/BiblioExchangeV2/core/di"
	"github.com/lastvoidtemplar/BiblioExchangeV2/core/di/identificators"
	"github.com/lastvoidtemplar/BiblioExchangeV2/core/server/validation"
	"github.com/lastvoidtemplar/BiblioExchangeV2/core/utils"
	"github.com/lastvoidtemplar/BiblioExchangeV2/library-service/dto"
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
			return utils.ErrorHandler(c, http.StatusInternalServerError, err)
		}
		if authors == nil {
			return c.JSON(http.StatusOK, []struct{}{})
		}
		return c.JSON(http.StatusOK, authors)
	}
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
			return utils.ErrorHandler(c, http.StatusInternalServerError, err)
		}
		if author == nil {
			return utils.ErrorHandler(c, http.StatusNotFound, "Author with this id isn`t found!")
		}

		if author.R == nil {
			return utils.ErrorHandler(c, http.StatusNotFound, "Author page ratings are missing!")
		}

		userId, anonymous, err := utils.GetUserId(c)

		if err != nil {
			return utils.ErrorHandler(c, http.StatusInternalServerError, err)
		}

		err = queries.AddAuthorPageView(ctx, db, id, userId, anonymous)

		if err != nil {
			log.Printf("Error: %s\n", err.Error())
			return utils.ErrorHandler(c, http.StatusInternalServerError,
				fmt.Sprintf("Error when adding page view for author with id %s!", id))
		}

		dto := dto.SingleAuthorDTO{
			Author:     *author,
			CountStars: len(author.R.Authorpageratings),
		}
		return c.JSON(http.StatusOK, dto)
	}
}

func CreateAuthor(c *di.Container) echo.HandlerFunc {
	db, err := di.GetService[*sql.DB](c, identificators.Database)
	if err != nil {
		log.Fatalln(err)
	}
	ctx := context.Background()

	return func(c echo.Context) error {
		userId, anonymous, err := utils.GetUserId(c)

		if err != nil {
			return utils.ErrorHandler(c, http.StatusInternalServerError, err)
		}

		if anonymous {
			return utils.ErrorHandler(c, http.StatusUnauthorized, "Must sign-in to create author page!")
		}
		var dto dto.AuthorCreateDTO

		if err := c.Bind(&dto); err != nil {
			return utils.ErrorHandler(c, http.StatusBadRequest, err)
		}

		if errors := dto.Valid(); !errors.Empty() {
			return validation.HandleValidationErrors(c, errors)
		}

		author := dto.Map()

		if err := queries.CreateAuthor(ctx, db, &author); err != nil {
			return utils.ErrorHandler(c, http.StatusInternalServerError, err)
		}

		if err := queries.AddAuthorPageView(ctx, db, author.AuthorID, userId, anonymous); err != nil {
			return utils.ErrorHandler(c, http.StatusInternalServerError, err)
		}

		return c.NoContent(http.StatusCreated)
	}
}
