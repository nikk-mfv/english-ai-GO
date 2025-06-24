package usecase

import (
	"context"
	"englishAI/repository"
)

type userUploadAvatarUsecase struct {
	userRepository repository.IUserRepository
}

func NewUserUploadAvatarUsecase(userRepository repository.IUserRepository) IUserUploadAvatarUsecase {
	return &userUploadAvatarUsecase{userRepository: userRepository}
}

func (uc *userUploadAvatarUsecase) Execute(ctx context.Context, userID uint, imageURL string) error {
	return uc.userRepository.UploadAvatar(ctx, userID, imageURL)
}
