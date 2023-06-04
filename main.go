package main

import (
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/jedib0t/go-pretty/v6/text"
	"github.com/joho/godotenv"
	"math/rand"
	"net/http"
	"os"
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
var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		//host := r.URL.String()
		//strings.Contains(host, )
		return true
	},
}

func logError(err string) {
	panic(text.FgRed.Sprintf("[ERROR] - %s", err))
}
func logSuccess(who string, data string) {
	fmt.Println(text.FgGreen.Sprintf("[%s] - %s", who, data))
}
func logWarning(who string, data string) {
	fmt.Println(text.FgYellow.Sprintf("[%s] - %s", who, data))
}

func checkError(err error) bool {
	if err != nil {
		logError(err.Error())
		return true
	}
	return false
}

func main() {
	err := godotenv.Load()
	checkError(err)

	http.HandleFunc("/ws", wsHandler)

	http.HandleFunc("/newchat", func(res http.ResponseWriter, req *http.Request) {
		rand.Seed(time.Now().UnixNano())
		rand.Shuffle(len(letters), func(i, j int) {
			letters[i], letters[j] = letters[j], letters[i]
		})
		newId := letters[:5]
		fmt.Println(string(newId))

		http.Redirect(res, req, "/"+string(newId), http.StatusSeeOther)
	})

	logSuccess("SERVER", "Trying to listen on :"+os.Getenv("PORT"))
	err = http.ListenAndServe(":"+os.Getenv("PORT"), nil)
	checkError(err)
}
