package config

import (
	"os"

	"github.com/SpringCare/sh-go-workshop/internal/config/flags"
)

func LoadEnvFile() (*flags.Configuration, error) {
	var appConf flags.Configuration

	appConf.Postgres.Database = os.Getenv("DB_NAME")
	appConf.Postgres.Driver = os.Getenv("DB_DRIVER")
	appConf.Postgres.Username = os.Getenv("DB_USERNAME")
	appConf.Postgres.Password = os.Getenv("DB_PASSWORD")
	appConf.Postgres.Host = os.Getenv("DB_HOST")
	appConf.Postgres.Port = os.Getenv("DB_PORT")
	appConf.Address = os.Getenv("ADDRESS")

	appConf.CloudinaryConfig.CloudinaryName = os.Getenv("CLOUDINARY_NAME")
	appConf.CloudinaryConfig.CloudinaryUrl = os.Getenv("CLOUDINARY_URL")
	appConf.CloudinaryConfig.CloudinaryApi = os.Getenv("CLOUDINARY_API")
	appConf.CloudinaryConfig.CloudinarySecret = os.Getenv("CLOUDINARY_SECRET")

	return &appConf, nil
}
