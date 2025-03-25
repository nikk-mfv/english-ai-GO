package entities

import "gorm.io/gorm"

type Vocabulary struct {
	gorm.Model
	Name          string   `gorm:"type:varchar(50);unique;not null" binding:"required" json:"name"`
	Definition    []string `gorm:"type:varchar(200);not null" binding:"required" json:"definition"`
	Example       string   `gorm:"type:text" json:"example"`
	Pronunciation string   `gorm:"type:text" json:"pronunciation"`
}
