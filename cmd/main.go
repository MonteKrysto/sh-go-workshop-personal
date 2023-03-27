package main

import (
	"fmt"

	"github.com/SpringCare/sh-go-workshop/internal/controllers"
	pdb "github.com/SpringCare/sh-go-workshop/internal/db"
	"github.com/SpringCare/sh-go-workshop/internal/middlewares"
	"github.com/SpringCare/sh-go-workshop/internal/repositories"
	"github.com/SpringCare/sh-go-workshop/internal/services"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type App struct {
	db gorm.DB
	// getUserIdFromToken func(c *gin.Context) string
}

type Controllers struct {
	AuthController       controllers.AuthController
	LoginController      controllers.LoginController
	ProfileController    controllers.ProfileController
	ImageController      controllers.ImageController
	SubscriberController controllers.SubscriberController
	CampaignController   controllers.CampaignController
}

// This is the main function that starts the app
func main() {
	// fmt.Println(runtime.NumCPU())
	fmt.Println("Booting up the app")
	app := App{}
	//
	// // We keep the main function small so call the Run function to set everything up
	if err := app.Run(); err != nil {
		panic(err)
	}
	fmt.Println("App started")
}

// The Run function sets up the app
// We use the *App receiver - (app *App), which is exactly like a method on a class, so we can access the db field
// We use a pointer receiver because we want to modify the db field and assign the db connection to the app instance
// Otherwise, we would just have a copy of the db connection and not be able to access it outside the function
func (app *App) Run() error {
	fmt.Println("In run")
	db := pdb.Init()
	app.db = *db

	// We now have access to the db connection in the app instance
	// We can use it to set up the API
	controllers := initAPI(db)
	initRoutes(controllers)

	return nil
}

// InitAPI We are using dependency injection
func initAPI(db *gorm.DB) Controllers {
	// The repositories are responsible for interacting with the database
	// so, we pass in the db connection to it to interact with the database
	authRepository := repositories.NewAuthRepository(db)
	loginRepository := repositories.NewLoginRepository(db)
	profileRepository := repositories.NewProfileRepository(db)
	imageRepository := repositories.NewImageRepository(db)
	subscriberRepository := repositories.NewSubscriberRepository(db)
	campaignRepository := repositories.NewCampaignRepository(db)

	// The services are responsible for business logic, it depends on the repository
	// so, we pass in the auth repository to it
	authService := services.NewAuthService(authRepository)
	loginService := services.NewLoginService(loginRepository)
	profileService := services.NewProfileService(profileRepository)
	imageService := services.NewImageService(imageRepository)
	subscriberService := services.NewSubscriberService(subscriberRepository)
	campaignService := services.NewCampaignService(campaignRepository)

	// The auth controller is responsible for handling requests and responses
	// so, we pass in the auth service to it
	authController := controllers.NewAuthController(*authService)
	loginController := controllers.NewLoginController(*loginService)
	profileController := controllers.NewProfileController(*profileService)
	imageController := controllers.NewImageController(*imageService)
	subscriberController := controllers.NewSubscriberController(*subscriberService)
	campaignController := controllers.NewCampaignController(*campaignService)

	// Return the controllers to use when we set up the routes
	return Controllers{
		AuthController:       *authController,
		LoginController:      *loginController,
		ProfileController:    *profileController,
		ImageController:      *imageController,
		SubscriberController: *subscriberController,
		CampaignController:   *campaignController,
	}
}

func initRoutes(controllers Controllers) {
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
	api.GET("profiles", controllers.ProfileController.GetAll)
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
