package entities

import "gorm.io/gorm"

type History struct {
	gorm.Model
	Name   string `gorm:"type:varchar(100);unique;not null" json:"name" binding:"required"`
	UserID uint   `gorm:"not null" json:"user_id"`
	User   User   `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE;"`
}
