package models

import "time"

type MenuItem struct {
	ID               uint `gorm:"primaryKey"`
	ModuleID         uint `gorm:"not null"`
	HubClientID      uint `gorm:"not null"`
	EntityRegisterID *uint
	Title            string `gorm:"type:jsonb;not null"`
	Icon             string `gorm:"size:100"`
	Type             string `gorm:"size:100"`
	Link             string
	MenuOrder        int
	ViewType         string `gorm:"size:255;not null;default:'public'"`
	ActiveOnHeader   bool   `gorm:"default:true"`
	ActiveOnMenu     bool   `gorm:"default:true"`
	ActiveOnFooter   bool   `gorm:"default:false"`
	IsDeletable      bool   `gorm:"default:true"`
	Active           bool   `gorm:"default:true"`
	CreatedAt        time.Time
	UpdatedAt        time.Time

	Module         Module          `gorm:"foreignKey:ModuleID"`
	HubClient      HubClient       `gorm:"foreignKey:HubClientID"`
	EntityRegister *EntityRegister `gorm:"foreignKey:EntityRegisterID"`
}
