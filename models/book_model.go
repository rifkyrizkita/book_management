package models

import "gorm.io/gorm"

type Book struct {
	gorm.Model
	Title       string   `gorm:"type:varchar(255);not null;unique"`
	Author      string   `gorm:"type:varchar(255);not null"`
	Synopsis    string   `gorm:"type:text;not null"`
	ISBN        string   `gorm:"type:varchar(255);not null;unique"`
	PublishedAt string   `gorm:"type:date;not null"`
	Publisher   string   `gorm:"type:varchar(255);not null"`
	Image       *string  `gorm:"type:varchar(255)"`
	Stock       uint     `gorm:"default:0"`
	CategoryID  uint     `gorm:"not null"`
	Category    Category `gorm:"foreignKey:CategoryID"`
}
