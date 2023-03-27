package controllers

import (
	"github.com/SpringCare/sh-go-workshop/internal/models"
	"github.com/SpringCare/sh-go-workshop/internal/services"
	"github.com/gin-gonic/gin"
)

type LoginController struct {
	LoginService services.LoginService
}

func NewLoginController(ls services.LoginService) *LoginController {
	return &LoginController{LoginService: ls}
}

func (lc *LoginController) Login(c *gin.Context) {
	// Get the user from the request body
	input := models.Login{}
	err := c.BindJSON(&input)
	if err != nil {
		c.JSON(500, gin.H{"message": err.Error()})
		return
	}

	// Call the service
	authedUser, _, authErr := lc.LoginService.ValidateLogin(input)
	if authErr != nil {
		c.JSON(500, gin.H{"message": authErr.Error()})
		return
	}

	// Return the response
	c.JSON(200, gin.H{"message": authedUser})
}
