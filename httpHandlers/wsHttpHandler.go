package httpHandlers

import (
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

requires `room` in url query

runs wsReader to handle WS queries
*/
func WsHttpHandler(res http.ResponseWriter, req *http.Request) {
	logs.LogSuccess("HTTP", "Got request on "+req.URL.String())

	//getting room id
	if !req.URL.Query().Has("room") {
		res.WriteHeader(http.StatusBadRequest)
		return
	}
	roomId := req.URL.Query().Get("room")

	// creating WS connection
	conn, err := upgrader.Upgrade(res, req, nil)
	if logs.CheckError(err) {
		res.WriteHeader(http.StatusInternalServerError)
		return
	}

	// creating new room if provided one doesn't exist
	if rooms[roomId] == nil {
		rooms[roomId] = &types.Room{
			Clients:        make(map[*websocket.Conn]bool),
			MessageChannel: make(chan types.WSMessage),
		}
	}
	rooms[roomId].Clients[conn] = true

	go wsHandlers.WsReader(conn, rooms[roomId])
}
