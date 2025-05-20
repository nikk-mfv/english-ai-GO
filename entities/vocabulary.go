package entities

import "gorm.io/gorm"

type Vocabulary struct {
	gorm.Model
	Name          string   `gorm:"type:varchar(50);unique;not null" json:"name"`
	Definition    string   `gorm:"type:text;not null" json:"definition"`
	Example       string   `gorm:"type:text" json:"example"`
	Pronunciation string   `gorm:"type:text" json:"pronunciation"`
	Topics        []*Topic `gorm:"many2many:topic_vocabularies;constraint:OnDelete:CASCADE"`
}
