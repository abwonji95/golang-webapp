package main

import (
	"io"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	gindump "github.com/tpkeeper/gin-dump"
	"mywebapp.com/m/controller"
	"mywebapp.com/m/middlewares"
	"mywebapp.com/m/service"
)

var videoService service.VideoService = service.New()
var videoController controller.VideoController = controller.New(videoService)

func setupLogOutput() {
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
}

func main() {
	setupLogOutput()
	server := gin.New()
	server.Use(gin.Recovery(), middlewares.Logger(), middlewares.BasicAuth(), gindump.Dump())

	server.GET("/test", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "Ok",
		})
	})

	server.GET("/videos", func(ctx *gin.Context) {
		ctx.JSON(200, videoController.FindAll())

	})

	server.POST("/videos", func(ctx *gin.Context) {
		err := videoController.Save(ctx)
		ctx.JSON(201, videoController.Save(ctx))
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}

	})
	server.Run(":8080")
}
