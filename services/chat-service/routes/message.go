package routes

import (
	"chat-service/dto"
	"context"
	"database/sql"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	dbmodels "github.com/lastvoidtemplar/BiblioExchangeV2/core/db/models"
	"github.com/lastvoidtemplar/BiblioExchangeV2/core/di"
	"github.com/lastvoidtemplar/BiblioExchangeV2/core/di/identificators"
	"github.com/lastvoidtemplar/BiblioExchangeV2/core/utils"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

func PostMessage(c *di.Container) echo.HandlerFunc {
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
			return utils.ErrorHandler(c, http.StatusUnauthorized, "Must sign-in to use the chat!")
		}

		var msg dto.Message
		if err := c.Bind(&msg); err != nil {
			return utils.ErrorHandler(c, http.StatusBadRequest, err)
		}

		var model dbmodels.Message
		model.Content = null.StringFrom(msg.Content)
		model.User1ID = userId
		model.User2ID = msg.UserID

		err = model.Insert(ctx, db, boil.Blacklist(dbmodels.MessageColumns.MessageID))
		BroadcastMessage(userId, msg)
		if err != nil {
			return utils.ErrorHandler(c, http.StatusInternalServerError, err)
		}

		return c.NoContent(http.StatusCreated)
	}
}

func GetMessagesMessage(c *di.Container) echo.HandlerFunc {
	db, err := di.GetService[*sql.DB](c, identificators.Database)
	if err != nil {
		log.Fatalln(err)
	}
	return func(c echo.Context) error {
		ctx := context.Background()
		id := c.Param("id")

		userId, anonymous, err := utils.GetUserId(c)

		if err != nil {
			return utils.ErrorHandler(c, http.StatusInternalServerError, err)
		}

		if anonymous {
			return utils.ErrorHandler(c, http.StatusUnauthorized, "Must sign-in to use the chat!")
		}

		msgs, err := dbmodels.Messages(
			qm.Where("user1_id = ?", id),
			qm.And("user2_id = ?", userId),
			qm.Or("user1_id = ?", userId),
			qm.And("user2_id = ?", id),
		).All(ctx, db)

		if err != nil {
			return utils.ErrorHandler(c, http.StatusInternalServerError, err)
		}

		return c.JSON(http.StatusOK, msgs)
	}
}
