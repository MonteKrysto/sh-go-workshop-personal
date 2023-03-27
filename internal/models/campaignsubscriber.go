package models

type CampaignSubscriber struct {
	BaseModelWithUUID

	Campaign     Campaign   `gorm:"foreignKey:CampaignID"`
	CampaignID   string     `gorm:"type:uuid;index:,unique,composite:unique_email_per_user"`
	Subscriber   Subscriber `gorm:"foreignKey:SubscriberID"`
	SubscriberID string     `gorm:"type:uuid;index:,unique,composite:unique_email_per_user"`
}