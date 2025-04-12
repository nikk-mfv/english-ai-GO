package entities

import "gorm.io/gorm"

type Conversation struct {
	gorm.Model
	Name     string    `gorm:"type:varchar(100);not null" json:"name" binding:"required"`
	UserID   uint      `gorm:"not null" json:"user_id" binding:"required"`
	User     User      `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Messages []Message `gorm:"foreignKey:ConversationID;constraint:OnDelete:CASCADE;"`
}
