package entities

import "gorm.io/gorm"

type Message struct {
	gorm.Model
	Message        string       `gorm:"type:text;not null" json:"message" binding:"required"`
	ConversationID uint         `gorm:"not null" json:"conversation_id"`
	Conversation   Conversation `gorm:"foreignKey:ConversationID;constraint:OnDelete:CASCADE;" json:"-"`
	UserID         uint         `gorm:"not null" json:"user_id"`
	User           User         `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE;" json:"-"`
}
