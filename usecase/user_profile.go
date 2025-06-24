package usecase

import (
	"context"
	"englishAI/entities"
	"englishAI/repository"
)

type userProfileUsecase struct {
	userRepository repository.IUserRepository
}

func NewUserProfileUsecase(userRepository repository.IUserRepository) IUserProfileUsecase {
	return &userProfileUsecase{userRepository: userRepository}
}

func (uc *userProfileUsecase) Execute(ctx context.Context, userId uint) (*entities.User, error) {
	user, err := uc.userRepository.GetProfile(ctx, userId)

	if err != nil {
		return &entities.User{}, err
	}

	return user, nil
}
