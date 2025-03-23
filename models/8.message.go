package models

import "gorm.io/gorm"

type Message struct {
	gorm.Model
	IsUser  bool   `gorm:"not null"` // true: User, false: AI
	Message string `gorm:"type:text;not null"`
}
