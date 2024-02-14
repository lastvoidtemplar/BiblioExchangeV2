package routes

import (
	"context"
	"database/sql"
	"log"
	"net/http"
	"strings"
	"upload-service/upload"

	"github.com/labstack/echo/v4"
	dbmodels "github.com/lastvoidtemplar/BiblioExchangeV2/core/db/models"
	"github.com/lastvoidtemplar/BiblioExchangeV2/core/di"
	"github.com/lastvoidtemplar/BiblioExchangeV2/core/di/identificators"
	"github.com/lastvoidtemplar/BiblioExchangeV2/core/utils"
)

func CreateListing(c *di.Container) echo.HandlerFunc {
	uploadServer, err := di.GetService[*upload.UploadServerServer](c, upload.UploadServerIdentificator)
	if err != nil {
		log.Fatalln(err)
	}
	db, err := di.GetService[*sql.DB](c, identificators.Database)
	if err != nil {
		log.Fatalln(err)
	}
	return func(c echo.Context) error {
		ctx := context.Background()
		id := c.Param("id")
		file, err := c.FormFile("file")
		if err != nil {
			return utils.ErrorHandler(c, http.StatusBadRequest, err)
		}
		userId, anonymous, err := utils.GetUserId(c)
		if err != nil {
			return utils.ErrorHandler(c, http.StatusInternalServerError, err)
		}

		if anonymous {
			return utils.ErrorHandler(c, http.StatusUnauthorized, "Must sign-in to upload file!")
		}

		exist, err := dbmodels.UploadfileExists(ctx, db, id)
		if err != nil {
			return utils.ErrorHandler(c, http.StatusInternalServerError, err)
		}
		if !exist {
			return utils.ErrorHandler(c, http.StatusNotFound, "Invalid url!")
		}

		record, err := dbmodels.FindUploadfile(ctx, db, id)

		if err != nil {
			return utils.ErrorHandler(c, http.StatusInternalServerError, err)
		}

		if userId != record.UserID.String {
			uploadServer.RemoveUrl(id)
			return utils.ErrorHandler(c, http.StatusForbidden, "Wrong user id!")
		}

		fileExt := strings.Split(file.Filename, ".")[1]
		match := false
		for _, ext := range record.Allowedfileformats {
			if ext == fileExt {
				match = true
				break
			}
		}

		if !match {
			uploadServer.RemoveUrl(id)
			return utils.ErrorHandler(c, http.StatusBadRequest, "File has wrong extensions")
		}

		if file.Size > int64(record.Maxsize.Int) {
			uploadServer.RemoveUrl(id)
			return utils.ErrorHandler(c, http.StatusRequestEntityTooLarge, "File is too large!")
		}

		src, err := file.Open()
		if err != nil {
			return utils.ErrorHandler(c, http.StatusInternalServerError, err)
		}
		defer src.Close()

		err = uploadServer.Upload(id, strings.Split(file.Filename, ".")[1], file.Size, src)

		if err != nil {
			return utils.ErrorHandler(c, http.StatusInternalServerError, err)
		}

		return c.NoContent(http.StatusOK)
	}
}
