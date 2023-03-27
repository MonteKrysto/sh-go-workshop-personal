package models

import "github.com/google/uuid"

type UserImage struct {
	BaseModelWithUUID

	ID       uuid.UUID `gorm:"type:uuid;primary_key"`
	UserId   uuid.UUID `gorm:"type:uuid;primary_key"`
	ImageId  string    `gorm:"varchar(100)"`
	ImageUrl string    `gorm:"varchar(100)"`
}
