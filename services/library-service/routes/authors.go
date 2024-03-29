package routes

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/lastvoidtemplar/BiblioExchangeV2/core/authentication"
	dbmodels "github.com/lastvoidtemplar/BiblioExchangeV2/core/db/models"
	"github.com/lastvoidtemplar/BiblioExchangeV2/core/db/queries"
	"github.com/lastvoidtemplar/BiblioExchangeV2/core/di"
	"github.com/lastvoidtemplar/BiblioExchangeV2/core/di/identificators"
	"github.com/lastvoidtemplar/BiblioExchangeV2/core/server/validation"
	"github.com/lastvoidtemplar/BiblioExchangeV2/core/utils"
	"github.com/lastvoidtemplar/BiblioExchangeV2/library-service/dto"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

func GetAuthorsPaginated(c *di.Container) echo.HandlerFunc {
	db, err := di.GetService[*sql.DB](c, identificators.Database)
	if err != nil {
		log.Fatalln(err)
	}
	return func(c echo.Context) error {
		ctx := context.Background()
		page := 1
		limit := 10
		pageStr := c.QueryParam("page")
		limitStr := c.QueryParam("limit")
		if parsePageInt, err := strconv.Atoi(pageStr); err == nil {
			page = parsePageInt
		}
		if parseLimitInt, err := strconv.Atoi(limitStr); err == nil {
			limit = parseLimitInt
		}
		authors, err := queries.GetAllAuthors(ctx, db, page, limit)
		if err != nil {
			return utils.ErrorHandler(c, http.StatusInternalServerError, err)
		}
		if authors == nil {
			return c.NoContent(http.StatusNotFound)
		}
		return c.JSON(http.StatusOK, dto.MapAuthorsAToMinimizedAuthorDTOs(authors))
	}
}

func GetAuthorById(c *di.Container) echo.HandlerFunc {
	db, err := di.GetService[*sql.DB](c, identificators.Database)
	if err != nil {
		log.Fatalln(err)
	}

	return func(c echo.Context) error {
		ctx := context.Background()
		id := c.Param("id")

		author, err := queries.GetAuthorById(ctx, db, id)
		if err != nil {
			if err == sql.ErrNoRows {
				return utils.ErrorHandler(c, http.StatusNotFound, "Author with this id isn`t found!")
			}
			return utils.ErrorHandler(c, http.StatusInternalServerError, err)
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
			return utils.ErrorHandler(c, http.StatusInternalServerError,
				fmt.Sprintf("Error when adding page view for author with id %s!", id))
		}

		dto := dto.MapAuthorAToSingleAuthorDTO(author)

		return c.JSON(http.StatusOK, dto)
	}
}

func CreateAuthor(c *di.Container) echo.HandlerFunc {
	db, err := di.GetService[*sql.DB](c, identificators.Database)
	if err != nil {
		log.Fatalln(err)
	}

	return func(c echo.Context) error {
		ctx := context.Background()
		userId, anonymous, err := utils.GetUserId(c)

		if err != nil {
			return utils.ErrorHandler(c, http.StatusInternalServerError, err)
		}

		if anonymous {
			return utils.ErrorHandler(c, http.StatusUnauthorized, "Must sign-in to create author page!")
		}

		var dto dto.AuthorBodyDTO

		if err := c.Bind(&dto); err != nil {
			return utils.ErrorHandler(c, http.StatusBadRequest, err)
		}

		if errors := dto.Valid(); !errors.Empty() {
			return validation.HandleValidationErrors(c, errors)
		}

		author := dto.Map()

		tx, err := db.BeginTx(ctx, &sql.TxOptions{})
		if err != nil {
			return utils.ErrorHandler(c, http.StatusInternalServerError, err)
		}
		defer tx.Rollback()

		if err := queries.CreateAuthor(ctx, tx, &author); err != nil {
			return utils.ErrorHandler(c, http.StatusInternalServerError, err)
		}

		if err := queries.AddAuthorPageView(ctx, tx, author.AuthorID, userId, anonymous); err != nil {
			return utils.ErrorHandler(c, http.StatusInternalServerError, err)
		}
		tx.Commit()
		return c.String(http.StatusCreated, fmt.Sprintf("/authors/%s", author.AuthorID))
	}
}

func UpdateAuthor(c *di.Container) echo.HandlerFunc {
	db, err := di.GetService[*sql.DB](c, identificators.Database)
	if err != nil {
		log.Fatalln(err)
	}

	return func(c echo.Context) error {
		ctx := context.Background()
		authorId := c.Param("id")
		_, anonymous, err := utils.GetUserId(c)

		if err != nil {
			return utils.ErrorHandler(c, http.StatusInternalServerError, err)
		}

		if anonymous {
			return utils.ErrorHandler(c, http.StatusUnauthorized, "Must sign-in to create author page!")
		}
		var inDto dto.AuthorBodyDTO

		if err := c.Bind(&inDto); err != nil {
			return utils.ErrorHandler(c, http.StatusBadRequest, err)
		}

		if errors := inDto.Valid(); !errors.Empty() {
			return validation.HandleValidationErrors(c, errors)
		}

		author := inDto.Map()
		author.AuthorID = authorId

		updatedAuthor, found, err := queries.UpdateAuthor(ctx, db, &author)
		if err != nil {
			return utils.ErrorHandler(c, http.StatusInternalServerError, err)
		}

		if !found {
			return utils.ErrorHandler(c, http.StatusNotFound, "Author with this id isn`t found!")
		}

		outDto := dto.MapAuthorAToSingleAuthorDTO(updatedAuthor)

		return c.JSON(http.StatusOK, outDto)
	}
}

func DeleteAuthor(c *di.Container) echo.HandlerFunc {
	db, err := di.GetService[*sql.DB](c, identificators.Database)
	if err != nil {
		log.Fatalln(err)
	}

	return func(c echo.Context) error {
		ctx := context.Background()
		authorId := c.Param("id")
		_, anonymous, err := utils.GetUserId(c)

		if anonymous {
			return utils.ErrorHandler(c, http.StatusUnauthorized, "Must sign-in to create author page!")
		}

		if err != nil {
			return utils.ErrorHandler(c, http.StatusInternalServerError, err)
		}

		roles, err := utils.GetUserRoles(c)

		if err != nil {
			return utils.ErrorHandler(c, http.StatusInternalServerError, err)
		}

		hasCorrectRole := false
		for _, role := range roles {
			if role == authentication.Admin {
				hasCorrectRole = true
				break
			}
		}

		if !hasCorrectRole {
			return utils.ErrorHandler(c, http.StatusForbidden, "Admin role is required to delete an author")
		}

		found, err := queries.DeleteAuthor(ctx, db, authorId)

		if err != nil {
			if err == queries.ErrAuthorHasBooks {
				return utils.ErrorHandler(c, http.StatusBadRequest, err)
			}
			return utils.ErrorHandler(c, http.StatusInternalServerError, err)
		}

		if !found {
			return utils.ErrorHandler(c, http.StatusNotFound, "Author with this id isn`t found!")
		}

		return c.NoContent(http.StatusNoContent)
	}
}

func ToggleAuthorStarRating(c *di.Container) echo.HandlerFunc {
	db, err := di.GetService[*sql.DB](c, identificators.Database)
	if err != nil {
		log.Fatalln(err)
	}

	return func(c echo.Context) error {
		ctx := context.Background()
		authorId := c.Param("id")
		userId, anonymous, err := utils.GetUserId(c)

		if err != nil {
			return utils.ErrorHandler(c, http.StatusInternalServerError, err)
		}

		if anonymous {
			return utils.ErrorHandler(c, http.StatusUnauthorized, "Must sign-in to star author page!")
		}

		starred, authorFound, err := queries.ToggleStarRatingOnAuthorPage(ctx, db, authorId, userId)

		if err != nil {
			return utils.ErrorHandler(c, http.StatusInternalServerError, err)
		}

		if !authorFound {
			return utils.ErrorHandler(c, http.StatusNotFound,
				fmt.Errorf("author with id %s was not found", authorId))
		}

		return c.JSON(http.StatusOK, dto.CreateStarDTO(starred))
	}
}

func CreateAuthorReview(c *di.Container) echo.HandlerFunc {
	db, err := di.GetService[*sql.DB](c, identificators.Database)
	if err != nil {
		log.Fatalln(err)
	}

	return func(c echo.Context) error {
		ctx := context.Background()
		authorId := c.Param("id")
		rootId := c.QueryParam("rootId")
		userId, anonymous, err := utils.GetUserId(c)

		if err != nil {
			return utils.ErrorHandler(c, http.StatusInternalServerError, err)
		}

		if anonymous {
			return utils.ErrorHandler(c, http.StatusUnauthorized, "Must sign-in to star author page!")
		}

		var dto dto.ReviewBodyDto

		if err := c.Bind(&dto); err != nil {
			return utils.ErrorHandler(c, http.StatusBadRequest, err)
		}

		if errors := dto.Valid(); !errors.Empty() {
			return validation.HandleValidationErrors(c, errors)
		}

		tx, err := db.BeginTx(ctx, &sql.TxOptions{})
		if err != nil {
			return utils.ErrorHandler(c, http.StatusInternalServerError, err)
		}
		defer tx.Rollback()
		if rootId != "" {
			reviewOnAuthorExist, err := queries.ExistAuthorReviewOnSpecificAuthor(ctx, tx, rootId, authorId)

			if err != nil {
				return utils.ErrorHandler(c, http.StatusInternalServerError, err)
			}

			if !reviewOnAuthorExist {
				return utils.ErrorHandler(c, http.StatusBadRequest, "The root review isn`t on that author!")
			}
		}
		reviewId, authorFound, rootFound, err := queries.CreateAuthorReview(ctx, tx, authorId, userId, dto.Content, rootId)

		if err != nil {
			return utils.ErrorHandler(c, http.StatusInternalServerError, err)
		}

		if !authorFound {
			return utils.ErrorHandler(c, http.StatusNotFound,
				fmt.Errorf("author with id %s was not found", authorId))
		}

		if !rootFound {
			return utils.ErrorHandler(c, http.StatusNotFound,
				fmt.Errorf("author review with id %s was not found", rootId))
		}
		tx.Commit()
		return c.String(http.StatusCreated, fmt.Sprintf("/authors/%s/review/%s", authorId, reviewId))
	}
}

func UpdateAuthorReview(c *di.Container) echo.HandlerFunc {
	db, err := di.GetService[*sql.DB](c, identificators.Database)
	if err != nil {
		log.Fatalln(err)
	}

	return func(c echo.Context) error {
		ctx := context.Background()
		reviewId := c.Param("id")
		userId, anonymous, err := utils.GetUserId(c)

		if err != nil {
			return utils.ErrorHandler(c, http.StatusInternalServerError, err)
		}

		if anonymous {
			return utils.ErrorHandler(c, http.StatusUnauthorized, "Must sign-in to star author page!")
		}

		var dto dto.ReviewBodyDto

		if err := c.Bind(&dto); err != nil {
			return utils.ErrorHandler(c, http.StatusBadRequest, err)
		}

		if errors := dto.Valid(); !errors.Empty() {
			return validation.HandleValidationErrors(c, errors)
		}

		tx, err := db.BeginTx(ctx, &sql.TxOptions{})

		if err != nil {
			return utils.ErrorHandler(c, http.StatusInternalServerError, err)
		}
		defer tx.Rollback()

		exist, err := dbmodels.AuthorreviewExists(ctx, tx, reviewId)

		if err != nil {
			return utils.ErrorHandler(c, http.StatusInternalServerError, err)
		}

		if !exist {
			return utils.ErrorHandler(c, http.StatusBadRequest, "Wrong author review id")
		}

		review, err := queries.GetAuthorReviewById(ctx, tx, reviewId)

		if err != nil {
			return utils.ErrorHandler(c, http.StatusInternalServerError, err)
		}

		if userId != review.UserID {
			return utils.ErrorHandler(c, http.StatusForbidden, "Only the creator of the review can edit the review!")
		}

		review.Content = null.StringFrom(dto.Content)
		_, err = review.Update(ctx, tx, boil.Infer())

		if err != nil {
			return utils.ErrorHandler(c, http.StatusInternalServerError, err)
		}

		tx.Commit()
		return c.NoContent(http.StatusNoContent)
	}
}

func DeleteAuthorReview(c *di.Container) echo.HandlerFunc {
	db, err := di.GetService[*sql.DB](c, identificators.Database)
	if err != nil {
		log.Fatalln(err)
	}
	return func(c echo.Context) error {
		ctx := context.Background()
		reviewId := c.Param("id")
		userId, anonymous, err := utils.GetUserId(c)

		if err != nil {
			return utils.ErrorHandler(c, http.StatusInternalServerError, err)
		}

		if anonymous {
			return utils.ErrorHandler(c, http.StatusUnauthorized, "Must sign-in to star author page!")
		}

		tx, err := db.BeginTx(ctx, &sql.TxOptions{})
		if err != nil {
			return utils.ErrorHandler(c, http.StatusInternalServerError, err)
		}
		defer tx.Rollback()

		exist, err := dbmodels.AuthorreviewExists(ctx, tx, reviewId)

		if err != nil {
			return utils.ErrorHandler(c, http.StatusInternalServerError, err)
		}

		if !exist {
			return utils.ErrorHandler(c, http.StatusBadRequest, "Wrong author review id")
		}

		review, err := queries.GetAuthorReviewById(ctx, tx, reviewId)

		if err != nil {
			return utils.ErrorHandler(c, http.StatusInternalServerError, err)
		}

		isAdmin := false
		roles, _ := utils.GetUserRoles(c)

		for _, role := range roles {
			if role == authentication.Admin {
				isAdmin = true
				break
			}
		}

		if !isAdmin && userId != review.UserID {
			return utils.ErrorHandler(c, http.StatusForbidden, "Only the creator of the review can edit the review!")
		}

		_, err = review.Delete(ctx, tx)

		if err != nil {
			return utils.ErrorHandler(c, http.StatusInternalServerError, err)
		}

		tx.Commit()
		return c.NoContent(http.StatusNoContent)
	}
}
