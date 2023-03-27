package routes

import (
	"fmt"

	"github.com/SpringCare/sh-go-workshop/internal/middlewares"
	"github.com/gin-gonic/gin"
)

type Controllers struct {
}

func NewRoutes(controllers Controllers) {

}

func InitRoutes(controllers Controllers) {
	fmt.Println("Setting up routes")
	// Now let's set up the router and define some routes
	router := gin.Default()

	// Let's check if the api is up and running
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// Routes pertaining to auths
	router.POST("register", controllers.AuthController.Auth)
	router.POST("login", controllers.LoginController.Login)

	// All routes below this line are protected and will require authentication by passing a token in the header
	api := router.Group("/api/v1")
	api.Use(middlewares.JwtAuthMiddleware())

	// Routes pertaining to profiles
	api.GET("profilesa", controllers.ProfileController.GetAll)
	api.GET("profiles/:id", controllers.ProfileController.GetById)
	api.POST("profiles", controllers.ProfileController.Create)
	api.PATCH("profiles/:id", controllers.ProfileController.Update)
	api.DELETE("profiles/:id", controllers.ProfileController.Delete)

	// Routes pertaining to images
	api.POST("/upload", controllers.ImageController.Upload)

	// Routes pertaining to users
	api.GET("/user", controllers.AuthController.CurrentUser)

	api.POST("/email-list", controllers.SubscriberController.Upload)

	api.POST("/campaign", controllers.CampaignController.Create)
	api.GET("/campaign", controllers.CampaignController.GetAll)
	api.GET("/campaign/:id", controllers.CampaignController.GetById)
	api.PATCH("/campaign/:id", controllers.CampaignController.Update)
	api.DELETE("/campaign/:id", controllers.CampaignController.Delete)

	err := router.Run()
	if err != nil {
		return
	}
}
