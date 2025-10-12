package controller

import (
	service "go/tutorial/Service"
	"go/tutorial/entity"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	service *service.UserService
}

func NewUserController(service *service.UserService) *UserController {
	return &UserController{service: service}
}

func (uc *UserController) FindAll(c *gin.Context) {
	users := uc.service.FindAll()
	c.JSON(http.StatusOK, users)
}

func (uc *UserController) FindById(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	user := uc.service.FindById(id)
	if user.Id == 0 {
		c.JSON(http.StatusNotFound, gin.H{"message": "User not found"})
		return
	}
	c.JSON(http.StatusOK, user)
}

func (uc *UserController) Save(c *gin.Context) {
	var user entity.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	saved := uc.service.Save(user)
	c.JSON(http.StatusCreated, saved)
}

func (uc *UserController) Delete(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	if uc.service.DeleteById(id) {
		c.JSON(http.StatusOK, gin.H{"message": "User deleted"})
	} else {
		c.JSON(http.StatusNotFound, gin.H{"message": "User not found"})
	}
}

func (uc *UserController) Update(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var user entity.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updated := uc.service.UpdateById(id, user)
	if updated.Id == 0 {
		c.JSON(http.StatusNotFound, gin.H{"message": "User not found"})
		return
	}
	c.JSON(http.StatusOK, updated)
}
