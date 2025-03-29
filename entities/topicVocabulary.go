package entities

import "time"

type TopicVocabulary struct {
	TopicID      uint       `gorm:"primaryKey" json:"topic_id"`
	VocabularyID uint       `gorm:"primaryKey" json:"vocabulary_id"`
	Topic        Topic      `gorm:"foreignKey:TopicID;constraint:OnDelete:CASCADE;" json:"-"`
	Vocabulary   Vocabulary `gorm:"foreignKey:VocabularyID;constraint:OnDelete:CASCADE;" json:"-"`
	CreatedAt    time.Time  `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt    time.Time  `gorm:"autoUpdateTime" json:"updated_at"`
}
