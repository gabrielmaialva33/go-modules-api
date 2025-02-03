package models

type Module struct {
	BaseID

	Title       string `gorm:"type:jsonb;not null"`
	Type        string `gorm:"size:50;not null"`
	Entities    int    `gorm:"default:1"`
	Unlimited   bool   `gorm:"default:true"`
	HubClientID uint   `gorm:"not null"`

	BaseAttributes
	BaseTimestamps

	HubClient HubClient `gorm:"foreignKey:HubClientID"`
}
