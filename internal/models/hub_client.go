package models

type HubClient struct {
	BaseID
	Name       string `gorm:"type:varchar(255);not null" json:"name"`
	Active     bool   `gorm:"default:true" json:"active"`
	ExternalID string `gorm:"unique;not null" json:"external_id"`
	BaseAttributes
	BaseTimestamps
}
