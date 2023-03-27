package repositories

import (
	"fmt"
	"mime/multipart"

	"github.com/SpringCare/sh-go-workshop/internal/config"
	"github.com/SpringCare/sh-go-workshop/internal/models"
	"github.com/SpringCare/sh-go-workshop/internal/viewmodels"
	"github.com/SpringCare/sh-go-workshop/pkg/token"
	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

type ImageRepository struct {
	db *gorm.DB
}

func NewImageRepository(db *gorm.DB) *ImageRepository {
	return &ImageRepository{db: db}
}

func (ir ImageRepository) Upload(fileName string, file multipart.File, ctx *gin.Context) (viewmodels.UserImage, error) {
	// Get the user id from the token
	id, err := token.ExtractTokenID(ctx)

	// Create a new instance of the viewmodel
	var userImageVm viewmodels.UserImage

	// Parse the id to a uuid
	userId, err := uuid.Parse(id)
	if err != nil {
		return userImageVm, err
	}
	fmt.Println("userId: ", userId)
	// Create our instance
	err = godotenv.Load("../.env")

	cfg, _ := config.LoadEnvFile()

	cld, _ := cloudinary.New()
	cld.Config.URL.Secure = true
	cld.Config.Cloud.CloudName = cfg.CloudinaryConfig.CloudinaryName
	cld.Config.Cloud.APIKey = cfg.CloudinaryConfig.CloudinaryApi
	cld.Config.Cloud.APISecret = cfg.CloudinaryConfig.CloudinarySecret

	resp, err := cld.Upload.Upload(ctx, file, uploader.UploadParams{
		PublicID:       fileName,
		UniqueFilename: api.Bool(false),
		Overwrite:      api.Bool(true)})

	userImage := models.UserImage{
		UserId:   userId,
		ImageId:  resp.PublicID,
		ImageUrl: resp.SecureURL,
	}

	// Save the image to the database
	err = ir.db.Create(&userImage).Error

	userImageVm = viewmodels.UserImage{
		ID:       userImage.ID,
		UserId:   userImage.UserId,
		ImageId:  userImage.ImageId,
		ImageUrl: userImage.ImageUrl,
	}
	return userImageVm, err
}
