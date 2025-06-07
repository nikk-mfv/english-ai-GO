package usecase

import (
	"context"
	"englishAI/entities"
	"englishAI/repository"
)

type userCreateUsecase struct {
	userRepository repository.IUserRepository
}

func NewUserCreateUsecase(userRepository repository.IUserRepository) IUserCreateUsecase {
	return &userCreateUsecase{userRepository: userRepository}
}

func (uc *userCreateUsecase) Execute(ctx context.Context, user entities.User) error {
	return uc.userRepository.Create(ctx, &user)
}
