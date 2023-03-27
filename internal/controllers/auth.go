package controllers

import (
	"github.com/SpringCare/sh-go-workshop/internal/models"
	"github.com/SpringCare/sh-go-workshop/internal/services"
	"github.com/SpringCare/sh-go-workshop/pkg/token"
	"github.com/gin-gonic/gin"
	"net/http"
)

type AuthController struct {
	AuthService services.AuthService
}

// NewAuthController This is a factory function that returns a new instance of the controller, think of it as a constructor
func NewAuthController(as services.AuthService) *AuthController {
	return &AuthController{AuthService: as}
}

func (ac *AuthController) Auth(c *gin.Context) {
	// Get the user from the request body
	user := models.User{}
	err := c.BindJSON(&user)
	if err != nil {
		c.JSON(500, gin.H{"message": err.Error()})
		return
	}

	// Call the service
	authedUser, authErr := ac.AuthService.Register(user)
	if authErr != nil {
		c.JSON(500, gin.H{"message": authErr.Error()})
		return
	}

	// Return the response
	c.JSON(200, gin.H{"message": authedUser})
}

func (ac *AuthController) CurrentUser(c *gin.Context) {

	userId, err := token.ExtractTokenID(c)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	u, err := ac.AuthService.GetCurrentUser(userId)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "success", "data": u})
}
