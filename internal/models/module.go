package models

import "time"

type Module struct {
	ID          uint   `gorm:"primaryKey"`
	HubClientID uint   `gorm:"not null"`
	Title       string `gorm:"type:jsonb;not null"`
	ModuleType  string `gorm:"size:50;not null"`
	Entities    int    `gorm:"default:1"`
	Unlimited   bool   `gorm:"default:true"`
	Active      bool   `gorm:"default:true"`
	CreatedAt   time.Time
	UpdatedAt   time.Time

	HubClient HubClient `gorm:"foreignKey:HubClientID"`
}
