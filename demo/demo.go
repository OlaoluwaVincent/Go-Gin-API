package demo

import (
	"fmt"
	service "go/tutorial/Service"
	"go/tutorial/entity"
)

func demo() {
	fmt.Println("Init project")

	videoService := service.NewVideoService()
	userService := service.NewUserService()

	userService.Save(entity.User{Id: 1, Name: "John Doe", Email: "john.doe@example.com", Age: 30})

	videoService.Save(entity.Video{Id: "1", Title: "First Video", Description: "This is the first video"})
	videoService.Save(entity.Video{Id: "2", Title: "Second Video", Description: "This is the second video"})

	// allVideos := videoService.FindAll()
	// allUsers := userService.FindAll()
	// fmt.Println("All Videos: ", allVideos)
	// fmt.Println("All Users: ", allUsers)

	user := userService.FindById(2)
	if user.Id <= 0 {
		fmt.Println("User not found")
	} else {
		fmt.Println("User found: ", user)
	}

}
