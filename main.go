package main

import (
	"net/http"

	"github.com/Cannedsans/Simple-go-chat/handlers"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.Static("/static", "./static")
	router.LoadHTMLFiles("html/index.html")

	router.GET("/", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "index.html", gin.H{})
	})

	router.GET("/ws", handlers.HandleWebSocket)

	go handlers.HandleMessages()
	router.Run()
}
