package models

type Subscriber struct {
	BaseModelWithUUID

	User      User   `gorm:"foreignKey:UserID"`
	UserID    string `gorm:"type:uuid;index:,unique,composite:unique_email_per_user"`
	FirstName string `gorm:"varchar(100)"`
	LastName  string `gorm:"varchar(100)"`
	Email     string `gorm:"varchar(100);index:,unique,composite:unique_email_per_user"`
	DOB       string `gorm:"varchar(10)"`
	Phone     string `gorm:"varchar(10)"`
	SSN       string `gorm:"varchar(4)"`
	City      string `gorm:"varchar(100)"`
	Country   string `gorm:"varchar(2)"`
	Interest1 string `gorm:"varchar(10)"`
	Interest2 string `gorm:"varchar(10)"`
	Interest3 string `gorm:"varchar(10)"`
}
