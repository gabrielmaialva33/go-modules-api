package models

import "time"

type EntityRegister struct {
	ID            uint   `gorm:"primaryKey"`
	StructureType string `gorm:"size:255;not null"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
}
