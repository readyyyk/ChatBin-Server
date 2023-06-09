package wsHandlers

import (
	"encoding/json"
	"github.com/gorilla/websocket"
	"github.com/readyyyk/chatbin-server/pkg/logs"
	"github.com/readyyyk/chatbin-server/pkg/types"
)

func MessageRequestHandler(dataJSON string, conn *websocket.Conn, room *types.Room) {
	var data types.Message
	err := json.Unmarshal([]byte(dataJSON), &data)
	logs.CheckError(err)

	data.Sender = room.Clients[conn]

	dataJSONwithSender, _ := json.Marshal(data)
	dataToWrite, _ := json.Marshal(types.FetchedDataS{
		Event: "message",
		Data:  string(dataJSONwithSender),
	})

	for clientConn := range room.Clients {
		if clientConn == conn {
			// in case we send back message event we send initial json
			// because we don't need to declare sender
			dataToWriteSameSender, _ := json.Marshal(types.FetchedDataS{
				Event: "message",
				Data:  dataJSON,
			})
			err = clientConn.WriteMessage(1, dataToWriteSameSender)
			logs.CheckError(err)
			continue
		}
		err = clientConn.WriteMessage(1, dataToWrite)
	}
}
