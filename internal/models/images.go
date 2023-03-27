package models

import "github.com/google/uuid"

type Images struct {
	BaseModelWithUUID

	UserId      uuid.UUID `gorm:"type:uuid;foreignKey:UserId;not null"`
	Url         string    `gorm:"varchar(100)"`
	IsMainImage bool      `gorm:"boolean;default:false"`
}
