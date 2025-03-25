package entities

import "gorm.io/gorm"

type Vocabulary struct {
	gorm.Model
	Name          string   `gorm:"type:varchar(50);unique;not null"`
	Definition    []string `gorm:"type:varchar(200);not null"`
	Example       string   `gorm:"type:text"`
	Pronunciation string   `gorm:"type:text"`
}
