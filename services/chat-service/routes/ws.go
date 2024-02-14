package routes

import (
	"chat-service/dto"
	"log"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
	"github.com/lastvoidtemplar/BiblioExchangeV2/core/di"
	"github.com/lastvoidtemplar/BiblioExchangeV2/core/utils"
)

var (
	connections = make(map[string]*websocket.Conn)
	mutex       sync.RWMutex
	upgrader    = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
)

func HandleWebSocket(c *di.Container) echo.HandlerFunc {
	return func(c echo.Context) error {

		userId, anonymous, err := utils.GetUserId(c)

		if err != nil {
			return utils.ErrorHandler(c, http.StatusInternalServerError, err)
		}

		if anonymous {
			return utils.ErrorHandler(c, http.StatusUnauthorized, "Must sign-in to use the chat!")
		}
		ws, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
		if err != nil {
			log.Println("upgrade:", err)
			return err
		}

		defer ws.Close()

		mutex.Lock()
		connections[userId] = ws
		mutex.Unlock()
		for {
			var msg dto.Message
			err := ws.ReadJSON(&msg)
			if err != nil {
				log.Println("read:", err)
				break
			}

			// broadcastMessage(userId, msg)
		}

		return nil
	}
}

func BroadcastMessage(userId string, msg dto.Message) {
	mutex.RLock()
	defer mutex.RUnlock()

	conn, ok := connections[msg.UserID]

	if !ok {
		return
	}

	err := conn.WriteJSON(dto.Message{
		UserID:  userId,
		Content: msg.Content,
	})

	if err != nil {
		log.Println("write:", err)
		delete(connections, userId)
		conn.Close()
	}
}
