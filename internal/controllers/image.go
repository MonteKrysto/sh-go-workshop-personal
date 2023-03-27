package controllers

import (
	"fmt"
	"github.com/SpringCare/sh-go-workshop/internal/services"
	"github.com/gin-gonic/gin"
)

type ImageController struct {
	ImageService services.ImageService
}

func NewImageController(imageService services.ImageService) *ImageController {
	return &ImageController{ImageService: imageService}
}

func (ic ImageController) Upload(c *gin.Context) {
	fmt.Println("Upload")
	// Call the service
	file, _, err := c.Request.FormFile("fileName")
	fmt.Println("File: ", file)
	if err != nil {
		c.JSON(500, gin.H{"message": err.Error()})
		return
	}
	result, err := ic.ImageService.Upload(c.PostForm("name"), file, c)
	if err != nil {
		c.JSON(500, gin.H{"message": err.Error()})
		return
	}

	// Return the response
	c.JSON(200, gin.H{"message": result})

	// Get the preferred name of the file if its not supplied

}
