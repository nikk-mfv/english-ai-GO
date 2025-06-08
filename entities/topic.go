package entities

import "gorm.io/gorm"

type Topic struct {
	gorm.Model
	Name         string        `gorm:"type:varchar(100);not null;uniqueIndex:idx_name_userid" json:"name"`
	UserID       uint          `gorm:"not null;uniqueIndex:idx_name_userid" json:"user_id"`
	User         *User         `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE" json:"-"`
	Vocabularies []*Vocabulary `gorm:"many2many:topic_vocabularies;constraint:OnDelete:CASCADE"`
}
