package main

import (
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/jedib0t/go-pretty/v6/text"
	"github.com/joho/godotenv"
	"math/rand"
	"net/http"
	"os"
	"strconv"
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
	fmt.Println(text.FgRed.Sprintf("[ERROR] - %s", err))
	panic(err)
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

func wsReader(conn *websocket.Conn) {
	for {
		event, data, err := conn.ReadMessage()
		if websocket.IsCloseError(err, 1001, 1005) {
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

func main() {
	err := godotenv.Load()
	checkError(err)

	http.HandleFunc("/ws", func(res http.ResponseWriter, req *http.Request) {
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
	})

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
