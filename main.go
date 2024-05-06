package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	PORT = ":8888"
)

func main() {
	var err error

	repository := NewRepository()
	controller := NewController(repository)

	server := gin.Default()
	server = setupRouter(server, controller)
	err = server.Run(PORT)
	if err != nil {
		log.Fatal(err)
	}
}

func setupRouter(router *gin.Engine, controller Controller) *gin.Engine {

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	router.POST("/api/v1/create", controller.CreatePostController)
	router.GET("/api/v1/posts", controller.ReadPostsController)

	return router
}
