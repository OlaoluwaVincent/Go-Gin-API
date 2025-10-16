package main

import (
	controller "go/tutorial/Controller"
	service "go/tutorial/Service"
	"go/tutorial/middlewares"
	"io"
	"os"

	"github.com/gin-gonic/gin"
	gindump "github.com/tpkeeper/gin-dump"
)

func setupLogOutput() {
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)

}

func main() {
	setupLogOutput()
	server := gin.Default()

	server.Use(middlewares.BasicAuth(), gindump.Dump())

	server.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	videoService := service.NewVideoService()
	videoController := controller.NewVideoController(videoService)

	server.GET("/videos", videoController.FindAll)
	server.POST("/videos", videoController.Save)
	server.DELETE("/videos/:id", videoController.Delete)

	userService := service.NewUserService()
	userController := controller.NewUserController(userService)

	// Step 3: Register routes
	server.GET("/users", userController.FindAll)
	server.GET("/users/:id", userController.FindById)
	server.POST("/users", userController.Save)
	server.PUT("/users/:id", userController.Update)
	server.DELETE("/users/:id", userController.Delete)

	// server.POST("/login", userController)

	server.Run(":8080")
}
