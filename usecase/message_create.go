package usecase

import (
	"englishAI/entities"
	"englishAI/repository"
)

type messageCreateUsecase struct {
	messageRepo repository.IMessageRepo
}

type IMessageCreateUseCase interface {
	Execute(newMessage *entities.Message) error
}

func NewMessageCreateUsecase(repo repository.IMessageRepo) IMessageCreateUseCase {
	return &messageCreateUsecase{
		messageRepo: repo,
	}
}

func (uc *messageCreateUsecase) Execute(newMessage *entities.Message) error {
	err := uc.messageRepo.Create(newMessage)
	if err != nil {
		return err
	}

	return nil
}
