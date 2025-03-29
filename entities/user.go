package entities

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string    `gorm:"type:varchar(50);unique;not null" binding:"required" json:"username"`
	Password string    `gorm:"type:varchar(255);not null" binding:"required" json:"-"`
	History  []History `gorm:"foreignKey:UserID" json:"history"`
	Topic    []Topic   `gorm:"foreignKey:UserID" json:"topic"`
}
