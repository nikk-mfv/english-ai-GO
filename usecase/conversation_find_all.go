package usecase

import (
	"englishAI/entities"
	"englishAI/repository"
)

type conversationFindAllUsecase struct {
	conversationRepo repository.IConversationRepository
}

func NewConversationFindAllUsecase(repo repository.IConversationRepository) ITopicFindAllUsecase {
	return &conversationFindAllUsecase{conversationRepo: repo}
}

func (uc *conversationFindAllUsecase) Execute(userId uint) ([]entities.Conversation, error) {
	conversations, err := uc.conversationRepo.FindAll(userId)
	return conversations, err
}
