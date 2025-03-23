package models

import "time"

type HistoryMessage struct {
	HistoryID uint      `gorm:"primaryKey"`
	MessageID uint      `gorm:"primaryKey"`
	History   History   `gorm:"foreignKey:HistoryID;constraint:OnDelete:CASCADE;"`
	Message   Message   `gorm:"foreignKey:MessageID;constraint:OnDelete:CASCADE;"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}
