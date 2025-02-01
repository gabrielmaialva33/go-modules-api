package models

import "time"

type ModuleEntityRegister struct {
	ID               uint `gorm:"primaryKey"`
	ModuleID         uint `gorm:"not null;index"`
	EntityRegisterID uint `gorm:"not null;index"`
	CreatedAt        time.Time
	UpdatedAt        time.Time

	Module         Module         `gorm:"foreignKey:ModuleID"`
	EntityRegister EntityRegister `gorm:"foreignKey:EntityRegisterID"`

	_ struct{} `gorm:"uniqueIndex:idx_module_entity"`
}
