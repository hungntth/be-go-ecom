package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/hungntth/be-go-ecom/internal/service"
	"github.com/hungntth/be-go-ecom/pkg/response"
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

	// response.SucessResponse(c, 20001, []string{"tips", "m10", "anoystick"})
	response.ErrorResponse(c, 20003, "Some thing is wrong")

}