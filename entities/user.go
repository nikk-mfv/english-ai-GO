package entities

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username      string `gorm:"type:varchar(50);unique;not null" binding:"required" json:"username"`
	Password      string `gorm:"type:varchar(255);not null" binding:"required" json:"password"`
	Email         string `gorm:"type:varchar(100);unique;default:null" json:"email"`
	Provider      string `gorm:"type:varchar(50);default:null" json:"provider,omitempty"`
	ProviderID    string `gorm:"type:varchar(100);default:null" json:"provider_id,omitempty"`
	ImageUrl      string `gorm:"type:varchar(255);default:null" json:"image_url,omitempty"`
	EmailImageUrl string `gorm:"type:varchar(255);default:null" json:"email_image_url,omitempty"`
}

func (u *User) GetAvatarURL() string {
	if u.ImageUrl != "" {
		return u.ImageUrl
	}
	if u.EmailImageUrl != "" {
		return u.EmailImageUrl
	}
	return ""
}
