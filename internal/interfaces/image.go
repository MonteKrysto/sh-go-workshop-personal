package interfaces

import (
	"github.com/SpringCare/sh-go-workshop/internal/viewmodels"
	"github.com/gin-gonic/gin"
	"mime/multipart"
)

type Image interface {
	Upload(fileName string, file multipart.File, ctx *gin.Context) (viewmodels.UserImage, error) // (models.Image, error)
}
