package main

import (
	"go-goroutine/controller"
	"runtime"

	"github.com/gin-gonic/gin"
)

func main() {
	runtime.GOMAXPROCS(1)

	router := gin.Default()
	router.GET("go1", controller.Goroutine1)
	router.GET("go2", controller.Gorutine2)
	router.Run(":8080")
}
