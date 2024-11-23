package models

import (
	"gorm.io/gorm"
)

type Borrowing struct {
	gorm.Model
	UserID     uint    `gorm:"not null"`
	User       User    `gorm:"foreignKey:UserID"`
	BookID     uint    `gorm:"not null"`
	Book       Book    `gorm:"foreignKey:BookID"`
	BorrowDate string  `gorm:"type:date;not null"`
	ReturnDate *string `gorm:"type:date;default:null"`
	Status     string  `gorm:"type:enum('borrowed','returned');default:'borrowed'"`
}
