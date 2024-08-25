package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hungntth/be-go-ecom/internal/service"
)

type UserController struct{
	UserService *service.UserService
}

func NewUserController() *UserController{
	return &UserController{
		UserService: service.NewUserService(),
	}
}

func (uc *UserController) GetUserById(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": uc.UserService.GetInfoUser(),
	})
}