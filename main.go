package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/readyyyk/chatbin-server/httpHandlers"
	"github.com/readyyyk/chatbin-server/pkg/logs"
	"math/rand"
	"net/http"
	"os"
	"time"
)

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

// Starting http server and http routes
func main() {
	err := godotenv.Load()
	logs.CheckError(err)

	http.HandleFunc("/ws", httpHandlers.WsHttpHandler)

	http.HandleFunc("/newchat", func(res http.ResponseWriter, req *http.Request) {
		rand.Seed(time.Now().UnixNano())
		rand.Shuffle(len(letters), func(i, j int) {
			letters[i], letters[j] = letters[j], letters[i]
		})
		newId := letters[:5]
		fmt.Println(string(newId))

		http.Redirect(res, req, "/"+string(newId), http.StatusSeeOther)
	})

	logs.LogSuccess("SERVER", "Trying to listen on :"+os.Getenv("PORT"))
	err = http.ListenAndServe(":"+os.Getenv("PORT"), nil)
	logs.CheckError(err)
}
