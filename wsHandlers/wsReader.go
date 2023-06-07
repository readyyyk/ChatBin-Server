package wsHandlers

import (
	"encoding/json"
	"github.com/gorilla/websocket"
	"github.com/readyyyk/chatbin-server/pkg/logs"
	"github.com/readyyyk/chatbin-server/pkg/types"
	"strconv"
)

var closeCodes = []int{websocket.CloseNormalClosure, websocket.CloseGoingAway, websocket.CloseNoStatusReceived, websocket.CloseAbnormalClosure}

/*
WsReader
 1. checks and validates initial data provided in WS query
 2. calls needed function based on `event` provided in query
*/
func WsReader(conn *websocket.Conn, room *types.Room) {
	for {
		//getting message from connection
		event, data, err := conn.ReadMessage()
		logs.LogSuccess("WS", "event: "+strconv.Itoa(event)+" "+string(data))

		// handle exit codes and errors
		if event == -1 || websocket.IsCloseError(err, closeCodes...) {
			logs.LogSuccess("WS", err.Error())
			return
		}
		if err != nil {
			logs.LogWarning("WS", err.Error())
		}

		// getting event name and data provided in a query
		// and handle wrong query format
		var fetchedData types.FetchedDataS
		err = json.Unmarshal(data, &fetchedData)
		if err != nil {
			messageBytes, _ := json.Marshal(types.ErrorMessage{
				Code:        404,
				Description: "Provide data in format {event: string, data: string}",
			})
			logs.CheckError(conn.WriteMessage(-1, messageBytes))
		}

		// calling needed function based on event name we got
		switch fetchedData.Event {
		case "connection":
			ConnectionRequestHandler(fetchedData.Data, conn)
			break
		default:
			// TODO: REPLACE WITH ERROR MESSAGE
			//room.MessageChannel <- types.WSMessage{
			//	Event: event,
			//	Data:  data,
			//}
			logs.LogWarning("WS", "No such event: "+fetchedData.Event)
			err = conn.WriteMessage(event, data)
			if logs.CheckError(err) {
				return
			}
			break
		}
	}
}
