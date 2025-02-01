package models

import "time"

type MenuItemPermission struct {
	ID         uint `gorm:"primaryKey"`
	RoleID     uint `gorm:"not null;index"`
	MenuItemID uint `gorm:"not null;index"`
	CreatedAt  time.Time
	UpdatedAt  time.Time

	Role     Role     `gorm:"foreignKey:RoleID"`
	MenuItem MenuItem `gorm:"foreignKey:MenuItemID"`

	_ struct{} `gorm:"uniqueIndex:idx_menuitem_role"`
}
