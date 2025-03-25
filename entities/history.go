package entities

import "gorm.io/gorm"

type History struct {
	gorm.Model
	Name   string `gorm:"type:varchar(100);unique;not null"`
	UserID uint   `gorm:"not null"`
	User   User   `gorm:"constraint:OnDelete:CASCADE;"`
}
