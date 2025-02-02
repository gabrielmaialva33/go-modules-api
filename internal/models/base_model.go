package models

import "time"

type BaseID struct {
	ID uint `gorm:"primaryKey" json:"id"`
}

type BaseAttributes struct {
	Active    bool `gorm:"default:true" json:"active"`
	IsDeleted bool `gorm:"default:false" json:"is_deleted"`
}

type BaseTimestamps struct {
	DeletedAt *time.Time `gorm:"index" json:"deleted_at"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
}

func (b BaseID) GetID() uint {
	return b.ID
}

// IDGetter is an interface for models that have an ID.
type IDGetter interface {
	GetID() uint
}
