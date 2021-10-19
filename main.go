package main

import (
	"go-goroutine/controller"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("go1", controller.Goroutine1)
	router.GET("go2", controller.Gorutine2)
	router.Run(":8080")
}
