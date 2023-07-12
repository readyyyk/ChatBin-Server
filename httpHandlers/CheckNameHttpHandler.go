package httpHandlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func CheckNameHttpHandler(c *gin.Context) {
	roomId := c.Param("chat")
	if roomId == "" {
		c.Status(http.StatusBadRequest)
		return
	}

	name := c.Query("name")

	for _, existingName := range rooms[roomId].Clients {
		if existingName == name {
			c.Status(http.StatusBadRequest)
			return
		}
	}

	c.Status(http.StatusOK)
	return
}
