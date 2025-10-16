package controller

import (
	service "go/tutorial/Service"
	"go/tutorial/entity"
	"go/tutorial/validators"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type VideoController struct {
	Service *service.VideoService
}

func (vc *VideoController) FindAll(c *gin.Context) {
	videos := vc.Service.FindAll()
	c.JSON(http.StatusOK, videos)
}

func (vc *VideoController) Save(c *gin.Context) {
	var video entity.Video
	if err := c.ShouldBindJSON(&video); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	savedVideo := vc.Service.Save(video)
	c.JSON(http.StatusOK, savedVideo)
}

func (vc *VideoController) Delete(c *gin.Context) {
	id := c.Param("id")
	if vc.Service.Delete(id) {
		c.JSON(http.StatusOK, gin.H{"message": "Deleted successfully"})
	} else {
		c.JSON(http.StatusNotFound, gin.H{"message": "Video not found"})
	}
}

var validate *validator.Validate

func NewVideoController(service *service.VideoService) *VideoController {
	validate = validator.New()
	validate.RegisterValidation("is-cool", validators.CoolValidator)
	return &VideoController{Service: service}
}
