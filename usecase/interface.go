package usecase

import (
	"context"
	"englishAI/entities"
)

type IVocabularyCreateUsecase interface {
	Execute(ctx context.Context, newVocabulary entities.Vocabulary) (entities.Vocabulary, error)
}

type IVocabularyFindUsecase interface {
	Execute(ctx context.Context) ([]entities.Vocabulary, error)
}

type IVocabularyDeleteUsecase interface {
	Execute(ctx context.Context, id string) error
}

type IVocabularyUpdateUsecase interface {
	Execute(ctx context.Context, id string, obj *entities.Vocabulary) error
}

// Conversation
type IConversationCreateUsecase interface {
	Execute(ctx context.Context, newConversation entities.Conversation) (entities.Conversation, error)
}
