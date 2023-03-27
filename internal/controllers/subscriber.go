package controllers

import (
	"fmt"
	"github.com/SpringCare/sh-go-workshop/internal/models"
	"github.com/SpringCare/sh-go-workshop/internal/services"
	"github.com/gin-gonic/gin"
)

type SubscriberController struct {
	SubscriberService services.SubscriberService
}

func NewSubscriberController(imageService services.SubscriberService) *SubscriberController {
	return &SubscriberController{SubscriberService: imageService}
}

func (ic SubscriberController) Upload(c *gin.Context) {
	fmt.Println("Upload")
	subscriber := models.Subscriber{}
	// Call the service
	file, err := c.FormFile("fileName")
	fmt.Println("File: ", file)
	if err != nil {
		c.JSON(500, gin.H{"message": err.Error()})
		return
	}

	err = c.SaveUploadedFile(file, "test.csv")
	err = ic.SubscriberService.Upload("test.csv", subscriber)
	if err != nil {
		c.JSON(500, gin.H{"message": err.Error()})
		return
	}

	// Return the response
	c.JSON(200, gin.H{"message": "done"})

	// Get the preferred name of the file if it's not supplied

}
