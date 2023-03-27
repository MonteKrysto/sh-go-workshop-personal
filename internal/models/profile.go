package models

import (
	"github.com/google/uuid"
)

type Profile struct {
	BaseModelWithUUID

	//ID      uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
	UserId  uuid.UUID `gorm:"type:uuid;foreignKey:UserId;not null" json:"user_id"`
	Street  string    `gorm:"varchar(100)"`
	City    string    `gorm:"varchar(100)"`
	State   string    `gorm:"varchar(2)"`
	ZipCode string    `gorm:"varchar(5)" json:"zip_code"`
}
