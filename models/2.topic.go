package models

import "gorm.io/gorm"

type Topic struct {
	gorm.Model
	Name string `gorm:"type:varchar(100);unique;not null"`
}
