package entities

import "time"

type TopicVocabulary struct {
	TopicID      uint       `gorm:"primaryKey"`
	VocabularyID uint       `gorm:"primaryKey"`
	Topic        Topic      `gorm:"constraint:OnDelete:CASCADE;"`
	Vocabulary   Vocabulary `gorm:"constraint:OnDelete:CASCADE;"`
	CreatedAt    time.Time  `gorm:"autoCreateTime"`
	UpdatedAt    time.Time  `gorm:"autoUpdateTime"`
}
