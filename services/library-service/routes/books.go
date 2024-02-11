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
	"github.com/volatiletech/sqlboiler/v4/boil"
)

func GetBooksPaginated(c *di.Container) echo.HandlerFunc {
	db, err := di.GetService[*sql.DB](c, identificators.Database)
	if err != nil {
		log.Fatalln(err)
	}

	return func(c echo.Context) error {
		ctx := context.Background()
		page := 1
		limit := 10
		authorId := c.Param("id")
		pageStr := c.QueryParam("page")
		limitStr := c.QueryParam("limit")
		if parsePageInt, err := strconv.Atoi(pageStr); err == nil {
			page = parsePageInt
		}
		if parseLimitInt, err := strconv.Atoi(limitStr); err == nil {
			limit = parseLimitInt
		}

		tx, err := db.BeginTx(ctx, &sql.TxOptions{})
		if err != nil {
			return utils.ErrorHandler(c, http.StatusInternalServerError, err)
		}
		defer tx.Rollback()

		exist, err := dbmodels.AuthorExists(ctx, tx, authorId)
		if err != nil {
			return utils.ErrorHandler(c, http.StatusInternalServerError, err)
		}
		if !exist {
			return utils.ErrorHandler(c, http.StatusBadRequest, "Wrong book id")
		}

		books, err := queries.GetAllBooks(ctx, tx, authorId, page, limit)

		if err != nil {
			return utils.ErrorHandler(c, http.StatusInternalServerError, err)
		}
		tx.Commit()
		return c.JSON(http.StatusOK, dto.MapBooksAToMinimizedBookDTOs(books))
	}
}

func GetBookById(c *di.Container) echo.HandlerFunc {
	db, err := di.GetService[*sql.DB](c, identificators.Database)
	if err != nil {
		log.Fatalln(err)
	}

	return func(c echo.Context) error {
		ctx := context.Background()
		id := c.Param("id")

		book, err := queries.GetBookById(ctx, db, id)
		if err != nil {
			if err == sql.ErrNoRows {
				return utils.ErrorHandler(c, http.StatusNotFound, "Book with this id isn`t found!")
			}
			return utils.ErrorHandler(c, http.StatusInternalServerError, err)
		}

		if book.R == nil {
			return utils.ErrorHandler(c, http.StatusNotFound, "Book page ratings are missing!")
		}

		userId, anonymous, err := utils.GetUserId(c)

		if err != nil {
			return utils.ErrorHandler(c, http.StatusInternalServerError, err)
		}

		err = queries.AddBookPageView(ctx, db, id, userId, anonymous)

		if err != nil {
			return utils.ErrorHandler(c, http.StatusInternalServerError,
				fmt.Sprintf("Error when adding page view for book with id %s!", id))
		}

		dto := dto.MapFromBook(book)

		return c.JSON(http.StatusOK, dto)
	}
}

func CreateBookById(c *di.Container) echo.HandlerFunc {
	db, err := di.GetService[*sql.DB](c, identificators.Database)
	if err != nil {
		log.Fatalln(err)
	}
	return func(c echo.Context) error {
		ctx := context.Background()

		userId, anonymous, err := utils.GetUserId(c)
		authorId := c.Param("id")

		if err != nil {
			return utils.ErrorHandler(c, http.StatusInternalServerError, err)
		}

		if anonymous {
			return utils.ErrorHandler(c, http.StatusUnauthorized, "Must sign-in to create book page!")
		}

		var dto dto.BookBodyDTO

		if err := c.Bind(&dto); err != nil {
			return utils.ErrorHandler(c, http.StatusBadRequest, err)
		}

		if errors := dto.Valid(); !errors.Empty() {
			return validation.HandleValidationErrors(c, errors)
		}

		book := dto.Map()

		tx, err := db.BeginTx(ctx, &sql.TxOptions{})
		if err != nil {
			return utils.ErrorHandler(c, http.StatusInternalServerError, err)
		}
		defer tx.Rollback()
		exist, err := dbmodels.AuthorExists(ctx, tx, authorId)
		if err != nil {
			return utils.ErrorHandler(c, http.StatusInternalServerError, err)
		}
		if !exist {
			return utils.ErrorHandler(c, http.StatusBadRequest, "Wrong book id")
		}

		if err := queries.CreateBook(ctx, tx, &book, authorId); err != nil {
			return utils.ErrorHandler(c, http.StatusInternalServerError, err)
		}
		if err := queries.AddBookPageView(ctx, tx, book.BookID, userId, anonymous); err != nil {
			return utils.ErrorHandler(c, http.StatusInternalServerError, err)
		}
		tx.Commit()
		return c.String(http.StatusCreated, fmt.Sprintf("/books/%s", book.BookID))
	}
}

func UpdateBook(c *di.Container) echo.HandlerFunc {
	db, err := di.GetService[*sql.DB](c, identificators.Database)
	if err != nil {
		log.Fatalln(err)
	}

	return func(c echo.Context) error {
		ctx := context.Background()
		bookId := c.Param("id")
		_, anonymous, err := utils.GetUserId(c)

		if err != nil {
			return utils.ErrorHandler(c, http.StatusInternalServerError, err)
		}

		if anonymous {
			return utils.ErrorHandler(c, http.StatusUnauthorized, "Must sign-in to create book page!")
		}
		var inDto dto.BookBodyDTO

		if err := c.Bind(&inDto); err != nil {
			return utils.ErrorHandler(c, http.StatusBadRequest, err)
		}

		if errors := inDto.Valid(); !errors.Empty() {
			return validation.HandleValidationErrors(c, errors)
		}

		book := inDto.Map()

		tx, err := db.BeginTx(ctx, &sql.TxOptions{})
		if err != nil {
			return utils.ErrorHandler(c, http.StatusInternalServerError, err)
		}
		defer tx.Rollback()

		exist, err := dbmodels.BookExists(ctx, tx, bookId)

		if err != nil {
			return utils.ErrorHandler(c, http.StatusInternalServerError, err)
		}

		if !exist {
			return utils.ErrorHandler(c, http.StatusNotFound, "Book with this id isn`t found!")
		}

		book.BookID = bookId

		updatedBook, err := queries.UpdateBook(ctx, tx, &book)
		if err != nil {
			return utils.ErrorHandler(c, http.StatusInternalServerError, err)
		}
		tx.Commit()
		outDto := dto.MapFromBook(updatedBook)

		return c.JSON(http.StatusOK, outDto)
	}
}

func DeleteBook(c *di.Container) echo.HandlerFunc {
	db, err := di.GetService[*sql.DB](c, identificators.Database)
	if err != nil {
		log.Fatalln(err)
	}

	return func(c echo.Context) error {
		ctx := context.Background()
		bookId := c.Param("id")
		_, anonymous, err := utils.GetUserId(c)

		if anonymous {
			return utils.ErrorHandler(c, http.StatusUnauthorized, "Must sign-in to create book page!")
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
			return utils.ErrorHandler(c, http.StatusForbidden, "Admin role is required to delete an book")
		}

		found, err := queries.DeleteBook(ctx, db, bookId)

		if err != nil {
			return utils.ErrorHandler(c, http.StatusInternalServerError, err)
		}

		if !found {
			return utils.ErrorHandler(c, http.StatusNotFound, "Book with this id isn`t found!")
		}

		return c.NoContent(http.StatusNoContent)
	}
}

func ToggleBookStarRating(c *di.Container) echo.HandlerFunc {
	db, err := di.GetService[*sql.DB](c, identificators.Database)
	if err != nil {
		log.Fatalln(err)
	}

	return func(c echo.Context) error {
		ctx := context.Background()
		bookId := c.Param("id")
		userId, anonymous, err := utils.GetUserId(c)

		if err != nil {
			return utils.ErrorHandler(c, http.StatusInternalServerError, err)
		}

		if anonymous {
			return utils.ErrorHandler(c, http.StatusUnauthorized, "Must sign-in to star book page!")
		}

		starred, bookFound, err := queries.ToggleStarRatingOnBookPage(ctx, db, bookId, userId)

		if err != nil {
			return utils.ErrorHandler(c, http.StatusInternalServerError, err)
		}

		if !bookFound {
			return utils.ErrorHandler(c, http.StatusNotFound,
				fmt.Errorf("book with id %s was not found", bookId))
		}

		return c.JSON(http.StatusOK, dto.CreateStarDTO(starred))
	}
}

func CreateBookReview(c *di.Container) echo.HandlerFunc {
	db, err := di.GetService[*sql.DB](c, identificators.Database)
	if err != nil {
		log.Fatalln(err)
	}

	return func(c echo.Context) error {
		ctx := context.Background()
		bookId := c.Param("id")
		rootId := c.QueryParam("rootId")
		userId, anonymous, err := utils.GetUserId(c)

		if err != nil {
			return utils.ErrorHandler(c, http.StatusInternalServerError, err)
		}

		if anonymous {
			return utils.ErrorHandler(c, http.StatusUnauthorized, "Must sign-in to star book page!")
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
			reviewOnBookExist, err := queries.ExistBookReviewOnSpecificBook(ctx, tx, rootId, bookId)

			if err != nil {
				return utils.ErrorHandler(c, http.StatusInternalServerError, err)
			}

			if !reviewOnBookExist {
				return utils.ErrorHandler(c, http.StatusBadRequest, "The root review isn`t on that book!")
			}
		}
		reviewId, bookFound, rootFound, err := queries.CreateBookReview(
			ctx, tx, bookId, userId, dto.Content, rootId)

		if err != nil {
			return utils.ErrorHandler(c, http.StatusInternalServerError, err)
		}

		if !bookFound {
			return utils.ErrorHandler(c, http.StatusNotFound,
				fmt.Errorf("book with id %s was not found", bookId))
		}

		if !rootFound {
			return utils.ErrorHandler(c, http.StatusNotFound,
				fmt.Errorf("book review with id %s was not found", rootId))
		}
		tx.Commit()
		return c.String(http.StatusCreated, fmt.Sprintf("/books/%s/review/%s", bookId, reviewId))
	}
}

func UpdateBookReview(c *di.Container) echo.HandlerFunc {
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
			return utils.ErrorHandler(c, http.StatusUnauthorized, "Must sign-in to star book page!")
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

		exist, err := dbmodels.BookreviewExists(ctx, tx, reviewId)

		if err != nil {
			return utils.ErrorHandler(c, http.StatusInternalServerError, err)
		}

		if !exist {
			return utils.ErrorHandler(c, http.StatusBadRequest, "Wrong book review id")
		}

		review, err := queries.GetBookReviewById(ctx, tx, reviewId)

		if err != nil {
			return utils.ErrorHandler(c, http.StatusInternalServerError, err)
		}

		if userId != review.UserID {
			return utils.ErrorHandler(c, http.StatusForbidden, "Only the creator of the review can edit the review!")
		}

		review.Content = dto.Content
		_, err = review.Update(ctx, tx, boil.Infer())

		if err != nil {
			return utils.ErrorHandler(c, http.StatusInternalServerError, err)
		}

		tx.Commit()
		return c.NoContent(http.StatusNoContent)
	}
}

func DeleteBookReview(c *di.Container) echo.HandlerFunc {
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
			return utils.ErrorHandler(c, http.StatusUnauthorized, "Must sign-in to star book page!")
		}

		tx, err := db.BeginTx(ctx, &sql.TxOptions{})
		if err != nil {
			return utils.ErrorHandler(c, http.StatusInternalServerError, err)
		}
		defer tx.Rollback()

		exist, err := dbmodels.BookreviewExists(ctx, tx, reviewId)

		if err != nil {
			return utils.ErrorHandler(c, http.StatusInternalServerError, err)
		}

		if !exist {
			return utils.ErrorHandler(c, http.StatusBadRequest, "Wrong book review id")
		}

		review, err := queries.GetBookReviewById(ctx, tx, reviewId)

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
