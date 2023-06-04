package main

import (
	"github.com/gorilla/websocket"
	"net/http"
	"strconv"
)

var closeCodes = []int{websocket.CloseNormalClosure, websocket.CloseGoingAway, websocket.CloseNoStatusReceived, websocket.CloseAbnormalClosure}

type queryData struct {
	Sender string `json:"sender"`
	Text   string `json:"text"`
	MType  string `json:"type"`
}
type wsQuery struct {
	Event string    `json:"event"`
	Data  queryData `json:"data"`
}

func wsReader(conn *websocket.Conn) {
	for {
		event, data, err := conn.ReadMessage()
		if websocket.IsCloseError(err, closeCodes...) {
			logSuccess("WS", err.Error())
			return
		}
		if checkError(err) {
			logWarning("WS", err.Error())
		}
		logSuccess("WS", "event: "+strconv.Itoa(event)+" "+string(data))
		err = conn.WriteMessage(event, data)
		if checkError(err) {
			return
		}
	}
}

func wsHandler(res http.ResponseWriter, req *http.Request) {
	logSuccess("HTTP", "Got request on "+req.URL.String())

	//getting room id
	if !req.URL.Query().Has("room") {
		res.WriteHeader(http.StatusBadRequest)
		return
	}
	roomId := req.URL.Query().Get("room")

	// creating WS connection
	conn, err := upgrader.Upgrade(res, req, nil)
	if checkError(err) {
		res.WriteHeader(http.StatusInternalServerError)
		return
	}

	if rooms[roomId] == nil {
		rooms[roomId] = &Room{
			Clients:        make(map[*websocket.Conn]bool),
			MessageChannel: make(chan Message),
		}
	}
	rooms[roomId].Clients[conn] = true

	go wsReader(conn)
}
