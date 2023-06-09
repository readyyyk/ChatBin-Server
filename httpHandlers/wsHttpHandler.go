package httpHandlers

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/readyyyk/chatbin-server/pkg/logs"
	"github.com/readyyyk/chatbin-server/pkg/types"
	"github.com/readyyyk/chatbin-server/wsHandlers"
	"net/http"
)

var rooms = make(map[string]*types.Room)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

/*
WsHttpHandler handles requests on /ws route
 1. requires `room` in url query
 2. runs wsReader to handle WS queries
*/
func WsHttpHandler(c *gin.Context) {
	//getting room id
	roomId := c.Param("chat")
	if roomId == "" {
		c.Status(http.StatusBadRequest)
		return
	}

	// creating WS connection
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if logs.CheckError(err) {
		c.Status(http.StatusInternalServerError)
		return
	}

	// creating new room if provided one doesn't exist
	if rooms[roomId] == nil {
		rooms[roomId] = &types.Room{
			Clients:        make(map[*websocket.Conn]string),
			MessageChannel: make(chan types.WSMessage),
		}
	}

	go wsHandlers.WsReader(conn, rooms[roomId])
}
