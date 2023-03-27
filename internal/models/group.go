package models

type Group struct {
	BaseModelWithUUID

	User   User   `gorm:"foreignKey:UserID"`
	UserID string `gorm:"type:uuid;"`
	Group  string `gorm:"varchar(100)"`
}