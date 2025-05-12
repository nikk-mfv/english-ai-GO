package usecase

import (
	"englishAI/entities"
	"englishAI/repository"
	"log"
)

type messageFindUsecase struct {
	messageRepo repository.IMessageRepo
}

type IMessageFindUseCase interface {
	Execute(conversationId string) ([]entities.Message, error)
}

func NewMessageFindUseCase(repo repository.IMessageRepo) IMessageFindUseCase {
	return &messageFindUsecase{
		messageRepo: repo,
	}
}

func (uc *messageFindUsecase) Execute(conversationId string) ([]entities.Message, error) {
	messages, err := uc.messageRepo.GetAll(conversationId)
	if err != nil {
		return []entities.Message{}, err
	}
	log.Print(messages)
	return messages, nil
}
