package models

import "github.com/google/uuid"

type Friends struct {
	BaseModelWithUUID

	UserId uuid.UUID `gorm:"type:uuid;foreignKey:UserId;not null"`
	Name   string    `gorm:"type:varchar(100);not null"`
}