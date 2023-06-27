package httpHandlers

import (
	"github.com/gin-gonic/gin"
	"github.com/readyyyk/chatbin-server/pkg/logs"
	"net/http"
)

func NameHttpHandler(c *gin.Context) {
	roomId := c.Param("chat")
	if roomId == "" {
		c.Status(http.StatusBadRequest)
		return
	}

	name := c.Query("name")

	for _, existingName := range rooms[roomId].Clients {
		logs.LogWarning("HTTP", existingName)
		if existingName == name {
			c.Status(http.StatusBadRequest)
			return
		}
	}

	c.Status(http.StatusOK)
	return
}
