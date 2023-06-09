package wsHandlers

import (
	"encoding/json"
	"github.com/gorilla/websocket"
	"github.com/readyyyk/chatbin-server/pkg/logs"
	"github.com/readyyyk/chatbin-server/pkg/types"
)

func ConnectionRequestHandler(dataJSON string, conn *websocket.Conn, room *types.Room) {
	var data types.ConnectionDataS
	err := json.Unmarshal([]byte(dataJSON), &data)
	logs.CheckError(err)

	if data.Detail == "connected" {
		if data.Name == "" {
			panic("data.Name is empty string in ConnectionRequestHandler")
		}
		room.Clients[conn] = data.Name
	} else if data.Detail == "disconnected" {
		delete(room.Clients, conn)
	}

	dataToWrite, _ := json.Marshal(types.FetchedDataS{
		Event: "connection",
		Data:  dataJSON,
	})

	for clientConn := range room.Clients {
		err = clientConn.WriteMessage(1, dataToWrite)
		logs.CheckError(err)
	}
}
