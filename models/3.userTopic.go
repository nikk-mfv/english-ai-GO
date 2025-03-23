package models

import "time"

type UserTopic struct {
	UserID    uint      `gorm:"primaryKey"`
	TopicID   uint      `gorm:"primaryKey"`
	User      User      `gorm:"constraint:OnDelete:CASCADE;"`
	Topic     Topic     `gorm:"constraint:OnDelete:CASCADE;"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}
