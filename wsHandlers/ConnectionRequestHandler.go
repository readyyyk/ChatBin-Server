package wsHandlers

import (
	"encoding/json"
	"github.com/gorilla/websocket"
	"github.com/readyyyk/chatbin-server/pkg/logs"
	"github.com/readyyyk/chatbin-server/pkg/types"
)

func ConnectionRequestHandler(dataJSON string, conn *websocket.Conn) {
	logs.LogWarning("WS", "Not implemented func for event `connection`")

	var data types.ConnectionDataS
	err := json.Unmarshal([]byte(dataJSON), &data)
	logs.CheckError(err)

	dataToWrite, _ := json.Marshal(types.FetchedDataS{
		Event: "connection",
		Data:  dataJSON,
	})
	err = conn.WriteMessage(1, dataToWrite)
	if logs.CheckError(err) {
		return
	}
}
