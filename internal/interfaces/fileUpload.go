package interfaces

import "github.com/SpringCare/sh-go-workshop/internal/models"

type FileUpload interface {
	Upload(fileName string, subscriber models.Subscriber) error
	//Create(profile models.Subscriber) error
}

//type CreateSubscriber interface {
//	Create(subscriber models.Subscriber) error
//}
