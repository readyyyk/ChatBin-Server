package main

import (
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/joho/godotenv"
	"github.com/readyyyk/terminal-todos-go/pkg/logs"
	"math/rand"
	"net/http"
	"time"
)

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

type Room struct {
	Clients        map[*websocket.Conn]bool
	MessageChannel chan Message
}
type Message struct {
	Room   string `json:"room"`
	Sender string `json:"sender"`
	Text   string `json:"text"`
	Type   string `json:"type"`
}

var rooms = make(map[string]*Room)

func main() {
	logs.LogError(godotenv.Load())

	http.HandleFunc("/ws", func(writer http.ResponseWriter, request *http.Request) {

	})
	http.HandleFunc("/newChat", func(writer http.ResponseWriter, request *http.Request) {
		rand.Seed(time.Now().UnixNano())
		rand.Shuffle(len(letters), func(i, j int) {
			letters[i], letters[j] = letters[j], letters[i]
		})
		newId := letters[:5]
		fmt.Println(string(newId))

		http.Redirect(writer, request, "/", http.StatusSeeOther)
	})
	logs.LogError(http.ListenAndServe(":3001", nil))
}
