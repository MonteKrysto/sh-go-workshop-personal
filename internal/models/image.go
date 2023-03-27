package models

import "github.com/google/uuid"

type Image struct {
	BaseModelWithUUID

	ID uuid.UUID `gorm:"type:uuid;primary_key"`
	Url string `gorm:"string" validate:"required"`

}
