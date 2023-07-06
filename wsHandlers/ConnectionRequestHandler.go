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
	if err != nil {
		dataToWrite, _ := json.Marshal(types.ErrorMessage{
			Code:        400,
			Description: "Invalid data in connection handler",
		})
		_ = conn.WriteMessage(websocket.TextMessage, dataToWrite)
	}

	if data.Detail == "connected" {
		if data.Name == "" {
			logs.LogError("data.Name is empty string in ConnectionRequestHandler")
		}
		room.Clients[conn] = data.Name
	} else if data.Detail == "disconnected" {
		delete(room.Clients, conn)
	} else if data.Detail == "trying to connect" {
		dataToWrite, _ := json.Marshal(types.FetchedDataS{
			Event: "connection",
			Data:  dataJSON,
		})

		for clientConn := range room.Clients {
			if clientConn == conn {
				continue
			}
			err = clientConn.WriteMessage(1, dataToWrite)
			if err != nil {
				logs.LogWarning("WS", err.Error())
			}
		}
	}
}
