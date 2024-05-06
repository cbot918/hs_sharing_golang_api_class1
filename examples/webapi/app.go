package main

import (
	"log"
	"net/http"

	// import module
	"github.com/gin-gonic/gin"
)

// PORT
const port = ":3001"

func main() {
	// 宣告 router engine
	r := gin.Default()

	// 宣告 API 路由
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	// router engine 監聽 PORT
	err := r.Run(port)
	if err != nil {
		log.Fatal(err)
	}
}
