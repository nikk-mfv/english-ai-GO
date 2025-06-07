package usecase

import (
	"context"
	"englishAI/entities"
	"englishAI/repository"
	"errors"

	"gorm.io/gorm"
)

type userFindUsecase struct {
	userRepository repository.IUserRepository
}

func NewUserFindUsecase(userRepository repository.IUserRepository) IUserFindUsecase {
	return &userFindUsecase{userRepository: userRepository}
}

func (uc *userFindUsecase) Execute(ctx context.Context, username string) (*entities.User, error) {
	user, err := uc.userRepository.FindByUsername(ctx, username)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		return &entities.User{}, err
	}

	return user, nil
}
