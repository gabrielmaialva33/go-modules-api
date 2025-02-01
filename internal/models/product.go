package models

import "time"

type Product struct {
	ID          uint   `gorm:"primaryKey"`
	Name        string `gorm:"type:jsonb;not null"`
	ProductType string `gorm:"size:255;not null"`
	IsDeleted   bool   `gorm:"default:false"`
	IsActive    bool   `gorm:"default:true"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
