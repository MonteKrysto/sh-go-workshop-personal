package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type BaseModelWithUUID struct {
	//gorm.Model
	ID        uuid.UUID      `gorm:"type:uuid;primary_key;"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"update_at"`
	DeletedAt gorm.DeletedAt `sql:"index" json:"deleted_at"`
}

// BeforeCreate is a hook that is called before a record is created
// This will generate a UUID for the ID field
func (model *BaseModelWithUUID) BeforeCreate(scope *gorm.DB) error {
	uuid := uuid.New()

	scope.Statement.SetColumn("ID", uuid)
	return nil
}
