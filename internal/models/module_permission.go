package models

import "time"

type ModulePermission struct {
	ID        uint `gorm:"primaryKey"`
	ModuleID  uint `gorm:"not null;index"`
	RoleID    uint `gorm:"not null;index"`
	CreatedAt time.Time
	UpdatedAt time.Time

	Module Module `gorm:"foreignKey:ModuleID"`
	Role   Role   `gorm:"foreignKey:RoleID"`

	_ struct{} `gorm:"uniqueIndex:idx_module_role"`
}
