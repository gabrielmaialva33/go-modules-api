package models

import "time"

type HubClient struct {
	ID         uint      `gorm:"primaryKey" json:"id"`
	Name       string    `gorm:"type:varchar(255);not null" json:"name"`
	Active     bool      `gorm:"default:true" json:"active"`
	ExternalID uint      `gorm:"unique;not null" json:"external_id"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}
