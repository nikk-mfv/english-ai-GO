package models

import "time"

type UserVocabulary struct {
	UserID       uint       `gorm:"primaryKey"`
	VocabularyID uint       `gorm:"primaryKey"`
	IsLearning   bool       `gorm:"not null"` //status true: "Learning", false: "mastered"
	User         User       `gorm:"constraint:OnDelete:CASCADE;"`
	Vocabulary   Vocabulary `gorm:"constraint:OnDelete:CASCADE;"`
	CreatedAt    time.Time  `gorm:"autoCreateTime"`
	UpdatedAt    time.Time  `gorm:"autoUpdateTime"`
}
