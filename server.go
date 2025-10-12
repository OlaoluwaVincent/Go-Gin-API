package main

import (
	controller "go/tutorial/Controller"
	service "go/tutorial/Service"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

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

	server.Run(":8080")
}
