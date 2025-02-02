package models

type Role struct {
	BaseID
	Name string `gorm:"type:varchar(50);not null" json:"name"`
	Slug string `gorm:"type:varchar(50);not null" json:"slug"`
	BaseAttributes
	BaseTimestamps
}
