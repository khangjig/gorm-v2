package socket

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
	"gorm-v2/model"
	"gorm-v2/util/myerror"
)

var u = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		// origin := r.Header.Get("Origin")
		// return origin == "http://localhost:3005"
		return true
	},
}

func (r *Route) socket(c echo.Context) error {
	var req = model.RequestSocket{}

	if err := c.Bind(&req); err != nil {
		return err
	}

	if len(req.Token) == 0 {
		// Handle Authentication here ...
		return myerror.ErrAuthentication(nil)
	}

	ws, err := u.Upgrade(c.Response(), c.Request(), nil)
	if err != nil {
		return err
	}

	defer ws.Close()

	for {
		// Received from Client
		_, msg, err := ws.ReadMessage()
		if err != nil {
			log.Fatal("Failed to ReadMessage: ", err.Error())
			c.Logger().Error(err)
		}

		//Handle received data
		receivedData := model.TestSocket{}

		err = json.Unmarshal(msg, &receivedData)
		if err != nil {
			log.Fatal("Failed to Unmarshal: ", err.Error())
		}

		req := model.TestSocket{
			To:      "client",
			Message: receivedData.Message,
		}

		data, err := json.Marshal(req)
		if err != nil {
			log.Fatal("Failed to Marshal: ", err.Error())
		}

		// Send to Client
		err = ws.WriteMessage(websocket.TextMessage, data)
		if err != nil {
			log.Fatal("Failed to WriteMessage: ", err.Error())
			c.Logger().Error(err)
		}
	}
}
