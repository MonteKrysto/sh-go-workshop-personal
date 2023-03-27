package models

import "github.com/google/uuid"

type Campaign struct {
	BaseModelWithUUID

	//User        User   `gorm:"foreignKey:UserID"`
	UserID      uuid.UUID `gorm:"type:uuid;foreignKey:UserID" json:"user_id"`
	Name        string    `gorm:"varchar(100)" json:"name"`
	Description string    `gorm:"varchar(100)" json:"description"`
	StartDate   string    `gorm:"varchar(10)" json:"start_date"`
	EndDate     string    `gorm:"varchar(10)" json:"end_date"`
	IsCompleted bool      `gorm:"varchar(10)"`
	IsDeleted   bool      `gorm:"varchar(10)"`
	IsFeatured  bool      `gorm:"varchar(10)"`
	IsSponsored bool      `gorm:"varchar(10)"`
	IsVerified  bool      `gorm:"varchar(10)"`
}
