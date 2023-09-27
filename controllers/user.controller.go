package controllers

import (
	"fmt"
	"net/http"
	"tecnotree-go-project/entities"
	"tecnotree-go-project/interfaces"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	UserService interfaces.IUser
}

func InitUserController(userSvc interfaces.IUser) *UserController {
	return &UserController{UserService: userSvc}
}

func (p UserController) HandleRegister(c *gin.Context) {

	fmt.Println("User controller Invoked")
	var register entities.Register
	err := c.BindJSON(&register)
	if err != nil {
		return
	}
	result, err := p.UserService.Register(&register)
	if err != nil {
		return
	} else {
		c.IndentedJSON(http.StatusCreated, result)
	}
}

func (p UserController) HandleLogin(c *gin.Context) {
	fmt.Println("User controller Invoked")
	var login entities.Login
	err := c.BindJSON(&login)
	if err != nil {
		return
	}
	result, err := p.UserService.Login(&login)
	if err != nil {
		return
	} else {
		c.IndentedJSON(http.StatusCreated, result)
	}
}

func (p UserController) HandleLogout(c *gin.Context) {
	result := p.UserService.Logout()
	c.IndentedJSON(http.StatusOK, result)
}
