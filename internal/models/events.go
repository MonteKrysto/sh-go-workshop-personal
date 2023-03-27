package models

import "github.com/google/uuid"

type Event struct {
	BaseModelWithUUID

	Name        string    `gorm:"string" validate:"required"`
	Description string    `gorm:"string" validate:"required"`
	Location    string    `gorm:"string" validate:"required"`
	StartTime   string    `gorm:"string" validate:"required"`
	ImageId     uuid.UUID `gorm:"type:uuid;foreignKey:ImageId;not null"`
}
