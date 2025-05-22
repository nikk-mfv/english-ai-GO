package entities

import "gorm.io/gorm"

type Vocabulary struct {
	gorm.Model
	Name          string   `gorm:"type:varchar(50);not null;uniqueIndex:idx_name_userid" json:"name"`
	Definition    string   `gorm:"type:text;not null" json:"definition"`
	Example       string   `gorm:"type:text" json:"example"`
	Pronunciation string   `gorm:"type:text" json:"pronunciation"`
	UserID        uint     `gorm:"not null;uniqueIndex:idx_name_userid" json:"user_id"`
	Topics        []*Topic `gorm:"many2many:topic_vocabularies;constraint:OnDelete:CASCADE"`
}
