package models

import "gorm.io/gorm"

type Category struct {
	gorm.Model
	Category string `gorm:"type:varchar(255);not null;unique"`
}
