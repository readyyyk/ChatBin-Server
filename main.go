package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/readyyyk/chatbin-server/httpHandlers"
	"github.com/readyyyk/chatbin-server/pkg/logs"
	"os"
)

// Starting gin server, http routes, cors policy
func main() {
	// err := godotenv.Load()
	// logs.CheckError(err)

	server := gin.Default()
	server.Use(cors.New(cors.Config{
		AllowAllOrigins:  true,
		AllowCredentials: true,
	}))

	server.GET("/:chat/ws", httpHandlers.WsHttpHandler)
	server.GET("/newchat", httpHandlers.NewchatHttpHandler)
	server.GET("/:chat/names", httpHandlers.NameHttpHandler)

	server.Static("/test", "./testPage")

	err := server.Run(":" + os.Getenv("PORT"))
	logs.CheckError(err)
}
