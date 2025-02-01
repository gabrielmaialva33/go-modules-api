package models

import "time"

type EntityRegisterProduct struct {
	ID               uint `gorm:"primaryKey"`
	ProductID        uint `gorm:"not null;index"`
	EntityRegisterID uint `gorm:"not null;index"`
	CreatedAt        time.Time
	UpdatedAt        time.Time

	Product        Product        `gorm:"foreignKey:ProductID"`
	EntityRegister EntityRegister `gorm:"foreignKey:EntityRegisterID"`

	_ struct{} `gorm:"uniqueIndex:idx_product_entity"`
}
