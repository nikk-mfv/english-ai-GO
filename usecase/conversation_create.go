package usecase

import (
	"context"
	"englishAI/entities"
	"englishAI/repository"
)

type conversationCreateUsecase struct {
	conversationRepository repository.IConversationRepository
}

func NewConversationCreateUsecase(conversationRepository repository.IConversationRepository) IConversationCreateUsecase {
	return &conversationCreateUsecase{conversationRepository: conversationRepository}
}

func (uc *conversationCreateUsecase) Execute(ctx context.Context, newConversation entities.Conversation) (entities.Conversation, error) {
	err := uc.conversationRepository.Create(ctx, &newConversation)
	if err != nil {
		return entities.Conversation{}, err
	}
	return newConversation, nil
}
