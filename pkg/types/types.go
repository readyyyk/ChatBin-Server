package types

import (
	"github.com/gorilla/websocket"
)

type Message struct {
	Sender string `json:"sender"`
	Text   string `json:"text"`
}
type Room struct {
	Clients        map[*websocket.Conn]string
	MessageChannel chan WSMessage
}

type WSMessage struct {
	Event int
	Data  []byte
}

type FetchedDataS struct {
	Event string `json:"event"`
	Data  string `json:"data"`
}

type ConnectionDataS struct {
	Detail string `json:"detail"` // "trying to connect", "connected", "disconnected"
	Name   string `json:"name"`
}
type MessageDataS struct {
	Sender string `json:"sender"`
	Text   string `json:"text"`
	MType  string `json:"type"`
}

type ErrorMessage struct {
	Code        int    `json:"code"` // http codes
	Description string `json:"description"`
}
