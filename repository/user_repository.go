package repository

import (
	"context"
	"englishAI/config"
	"englishAI/entities"
)

type userRepository struct{}

func NewUserRepository() IUserRepository {
	return &userRepository{}
}

func (r *userRepository) Create(ctx context.Context, newUser *entities.User) error {
	db := config.GetDatabase()

	return db.Create(newUser).Error
}

func (r *userRepository) FindByUsername(ctx context.Context, username string) (*entities.User, error) {
	var user entities.User
	db := config.GetDatabase()

	if err := db.Where("username = ?", username).First(&user).Error; err != nil {
		return &entities.User{}, err
	}

	return &user, nil
}
