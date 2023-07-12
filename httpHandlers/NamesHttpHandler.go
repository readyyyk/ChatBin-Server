package httpHandlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func NamesHttpHandler(c *gin.Context) {
	roomId := c.Param("chat")
	if roomId == "" {
		c.Status(http.StatusBadRequest)
		return
	}

	var nameList []string

	for _, name := range rooms[roomId].Clients {
		nameList = append(nameList, name)
	}

	fmt.Println(nameList)
	c.JSON(http.StatusOK, nameList)
}
