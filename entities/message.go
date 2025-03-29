package entities

import "gorm.io/gorm"

type Message struct {
	gorm.Model
	IsUser    bool    `gorm:"not null" json:"is_user" binding:"required"` // true: User, false: AI
	Message   string  `gorm:"type:text;not null" json:"message" binding:"required"`
	HistoryID uint    `gorm:"not null" json:"history_id"`
	History   History `gorm:"foreignKey:HistoryID;constraint:OnDelete:CASCADE;" json:"-"`
}
