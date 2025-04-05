package entities

import "gorm.io/gorm"

type Topic struct {
	gorm.Model
	Name         string        `gorm:"type:varchar(100);unique;not null" json:"name"`
	UserID       uint          `gorm:"not null" json:"user_id"`
	User         *User         `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE" json:"-"`
	Vocubularies []*Vocabulary `gorm:"many2many:topic_vocubularies;"`
}
