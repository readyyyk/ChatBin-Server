package httpHandlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"math/rand"
	"net/http"
	"time"
)

func NewchatHttpHandler(c *gin.Context) {
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(letters), func(i, j int) {
		letters[i], letters[j] = letters[j], letters[i]
	})
	newId := letters[:5]
	fmt.Println(string(newId))

	c.Redirect(http.StatusSeeOther, "/"+string(newId))
}
