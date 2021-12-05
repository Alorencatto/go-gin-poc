package main

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/augusto/golang-gin-poc/controller"
	"github.com/augusto/golang-gin-poc/midllewares"
	"github.com/augusto/golang-gin-poc/service"
	"github.com/gin-gonic/gin"
	gindump "github.com/tpkeeper/gin-dump"
)

var (
	videoService    service.VideoService       = service.New()
	videoController controller.VideoController = controller.New(videoService)
)

func setupLogOutput() {
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
}

func main() {

	setupLogOutput()

	server := gin.New()

	// server.Use(gin.Recovery())
	server.Use(gin.Recovery(), midllewares.Logger(), midllewares.BasicAuth(), gindump.Dump())

	server.Use()

	server.GET("/test", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "OK",
		})
	})

	server.GET("/videos", func(ctx *gin.Context) {
		ctx.JSON(200, videoController.FindAll())
	})

	server.POST("/videos", func(ctx *gin.Context) {
		err := videoController.Save(ctx)

		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		} else {
			ctx.JSON(http.StatusOK, gin.H{"message": "Video input is valid!"})
		}
	})

	server.Run(":8080")
	fmt.Println("Hello World")
}
