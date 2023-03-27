package services

import (
	"github.com/SpringCare/sh-go-workshop/internal/interfaces"
	"github.com/SpringCare/sh-go-workshop/internal/viewmodels"
	"github.com/gin-gonic/gin"
	"mime/multipart"
)

type ImageService struct {
	repository interfaces.Image
}

func NewImageService(repository interfaces.Image) *ImageService {
	return &ImageService{
		repository: repository,
	}
}

func (is ImageService) Upload(fileName string, file multipart.File, ctx *gin.Context) (viewmodels.UserImage, error) {
	return is.repository.Upload(fileName, file, ctx)
}
