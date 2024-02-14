package routes

import (
	"context"
	"database/sql"
	"listings-service/dto"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
	dbmodels "github.com/lastvoidtemplar/BiblioExchangeV2/core/db/models"
	"github.com/lastvoidtemplar/BiblioExchangeV2/core/db/queries"
	"github.com/lastvoidtemplar/BiblioExchangeV2/core/di"
	"github.com/lastvoidtemplar/BiblioExchangeV2/core/di/identificators"
	"github.com/lastvoidtemplar/BiblioExchangeV2/core/server/validation"
	"github.com/lastvoidtemplar/BiblioExchangeV2/core/upload"
	"github.com/lastvoidtemplar/BiblioExchangeV2/core/utils"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

var uploadTemplate = upload.NewUploadFileTemplate().
	SetAllowedFileFormats([]string{"png", "jpeg", "jpg"}).
	SetMaxSize(200 * 1024 * 1024)

func CreateListing(c *di.Container) echo.HandlerFunc {
	db, err := di.GetService[*sql.DB](c, identificators.Database)
	if err != nil {
		log.Fatalln(err)
	}
	uploadService, err := di.GetService[*upload.UploadService](c, identificators.UploadService)
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
			return utils.ErrorHandler(c, http.StatusUnauthorized, "Must sign-in to create book page!")
		}

		var dto dto.ListingBodyDto

		if err := c.Bind(&dto); err != nil {
			return utils.ErrorHandler(c, http.StatusBadRequest, err)
		}

		if errors := dto.Valid(); !errors.Empty() {
			return validation.HandleValidationErrors(c, errors)
		}

		listing := dto.Map()
		listing.OwnerID = userId

		tx, err := db.BeginTx(ctx, &sql.TxOptions{})
		if err != nil {
			return utils.ErrorHandler(c, http.StatusInternalServerError, err)
		}
		defer tx.Rollback()

		filesIds, err := queries.CreateListings(ctx, tx, &listing, dto.BooksIds, dto.NumberOfFiles)
		if err != nil {
			return utils.ErrorHandler(c, http.StatusInternalServerError, err)
		}

		urls := make([]string, 0, len(filesIds))

		for _, id := range filesIds {
			exp := time.Now().Add(2 * time.Minute)
			id := id
			t := uploadTemplate.SetPresignedUrlExpirationTime(&exp).
				SetPermission(userId).SetIdentificator(id).
				OnSuccessCallback(func(fileId string, url string) {
					log.Printf("On succ %s", fileId)
					file, err := dbmodels.FindListingsurl(ctx, db, id)
					if err != nil {
						return
					}
					file.ResourseURL = url
					file.UploadStatus = null.BoolFrom(true)

					file.Update(ctx, db, boil.Infer())
				}).
				OnErrorCallback(func(fileId string) {
					log.Printf("On err %s", fileId)
					file, err := dbmodels.FindListingsurl(ctx, db, id)
					if err != nil {
						return
					}
					file.UploadStatus = null.BoolFrom(true)

					file.Delete(ctx, db)
				}).
				OnUrlExpirationCallback(func(fileId string) {
					log.Printf("On expr %s", fileId)
					file, err := dbmodels.FindListingsurl(ctx, db, id)
					if err != nil {
						return
					}
					file.UploadStatus = null.BoolFrom(true)

					file.Delete(ctx, db)
				})
			url, err := uploadService.GetFilePresignedUrl(ctx, t)
			if err != nil {
				log.Println(err)
			}
			urls = append(urls, url)
		}

		tx.Commit()
		return c.JSON(http.StatusCreated, urls)
	}
}

func GetBooksPaginated(c *di.Container) echo.HandlerFunc {
	db, err := di.GetService[*sql.DB](c, identificators.Database)
	if err != nil {
		log.Fatalln(err)
	}

	return func(c echo.Context) error {
		ctx := context.Background()
		page := 1
		limit := 10
		bookId := c.Param("id")
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

		exist, err := dbmodels.BookExists(ctx, tx, bookId)
		if err != nil {
			return utils.ErrorHandler(c, http.StatusInternalServerError, err)
		}
		if !exist {
			return utils.ErrorHandler(c, http.StatusBadRequest, "Wrong book id")
		}

		listings, err := queries.GetAllListings(ctx, tx, bookId, page, limit)

		if err != nil {
			return utils.ErrorHandler(c, http.StatusInternalServerError, err)
		}
		tx.Commit()
		return c.JSON(http.StatusOK, dto.MapListingsAToListingSDisplayDtos(listings))
	}
}
