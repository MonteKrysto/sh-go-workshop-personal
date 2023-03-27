package repositories

import (
	"github.com/SpringCare/sh-go-workshop/internal/models"
	"gorm.io/gorm"
)

type SubscriberRepository struct {
	db *gorm.DB
}

func NewSubscriberRepository(db *gorm.DB) *SubscriberRepository {
	return &SubscriberRepository{db: db}
}

func (sr SubscriberRepository) Upload(fileName string, subscriber models.Subscriber) error {
	// Get the user id from the token
	//id, err := token.ExtractTokenID(ctx)
	//
	//// Parse the id to a uuid
	//userId, err := uuid.Parse(id)
	//if err != nil {
	//	return err
	//}
	//fmt.Println("userId: ", userId)

	//subscriber := models.Subscriber{
	//	FirstName: subscriber.FirstName,
	//	LastName:  subscriber.LastName,
	//	Email:     subscriber.Email,
	//	DOB:       subscriber.DOB,
	//	Phone:     subscriber.Phone,
	//	SSN:       subscriber.SSN,
	//	City:      subscriber.City,
	//	Country:   subscriber.Country,
	//	Interest1: subscriber.Interest1,
	//	Interest2: subscriber.Interest2,
	//	Interest3: subscriber.Interest3,
	//}

	//fmt.Println("subscriber: ", subscriber)

	// Save the image to the database
	err := sr.db.Create(&subscriber).Error

	return err
}
