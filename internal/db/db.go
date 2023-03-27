package db

import (
	"fmt"
	"log"
	"os"

	"github.com/SpringCare/sh-go-workshop/internal/models"

	"github.com/SpringCare/sh-go-workshop/internal/config"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func Init() *gorm.DB {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal("error: ", err)
	}

	cfg, _ := config.LoadEnvFile()

	connectString := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.Postgres.Host,
		cfg.Postgres.Port,
		cfg.Postgres.Username,
		cfg.Postgres.Password,
		cfg.Postgres.Database,
	)
	fmt.Println("string: " + connectString)

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			LogLevel: logger.Info, // Log level Info will output everything
		},
	)

	db, err := gorm.Open(postgres.Open(connectString), &gorm.Config{Logger: newLogger})
	if err != nil {
		log.Fatalf("Could not create postgres connection, err=%v", err)
		return db
	}

	db.AutoMigrate(
		&models.User{},
		&models.Profile{},
		&models.Image{},
		&models.UserImage{},
		&models.Subscriber{},
		&models.Group{},
		&models.Campaign{},
		&models.CampaignSubscriber{},
	)

	return db
}
