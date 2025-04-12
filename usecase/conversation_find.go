package usecase

import (
	"englishAI/entities"
	"englishAI/repository"
)

type conversationFindUsecase struct {
	conversationRepository repository.IConversationRepository
}

func NewConversationFindUseCase(repo repository.IConversationRepository) IConversationFindUsecase {
	return &conversationFindUsecase{conversationRepository: repo}
}

func (uc *conversationFindUsecase) Execute(id string) (entities.Conversation, error) {
	conversation, err := uc.conversationRepository.Find(id)

	return conversation, err
}
