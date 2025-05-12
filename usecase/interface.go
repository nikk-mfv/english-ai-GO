package usecase

import (
	"context"
	"englishAI/entities"
)

type IVocabularyCreateUsecase interface {
	Execute(ctx context.Context, newVocabulary entities.Vocabulary) (entities.Vocabulary, error)
}

type IVocabularyFindUsecase interface {
	Execute(ctx context.Context, paging entities.PagingRequest) ([]entities.Vocabulary, uint32, error)
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

type IConversationFindUsecase interface {
	Execute(id string) (entities.Conversation, error)
}

// Topic
type ITopicCreateUsecase interface {
	Execute(ctx context.Context, topic entities.Topic) (entities.Topic, error)
}

type ITopicFindUsecase interface {
	Execute(ctx context.Context, paging entities.PagingRequest) ([]entities.Topic, uint32, error)
}

type ITopicFindAllUsecase interface {
	Execute(userId int) ([]entities.Conversation, error)
}
