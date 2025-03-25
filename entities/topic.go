package entities

import "gorm.io/gorm"

type Topic struct {
	gorm.Model
	Name   string `gorm:"type:varchar(100);unique;not null"`
	UserID uint   `gorm:"not null"`
	User   User   `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE;"`
}
