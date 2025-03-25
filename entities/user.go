package entities

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string    `gorm:"type:varchar(50);unique;not null"`
	Password string    `gorm:"type:varchar(255);not null"`
	History  []History `gorm:"foreignKey:UserID"`
	Topic    []Topic   `gorm:"foreignKey:UserID"`
}
