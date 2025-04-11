package repository

import (
	"context"
	"englishAI/entities"
)

type IVocabularyRepository interface {
	Create(ctx context.Context, obj *entities.Vocabulary) error
	GetAll(ctx context.Context) ([]entities.Vocabulary, error)
	DeleteByID(ctx context.Context, id string) error
	UpdateByID(ctx context.Context, id string, obj *entities.Vocabulary) error
}

type IConversationRepository interface {
	Create(ctx context.Context, obj *entities.Conversation) error
}

type ITopicRepository interface {
	Create(ctx context.Context, topic *entities.Topic) error
	GetByPage(ctx context.Context, paging entities.PagingRequest) ([]entities.Topic, int64, error)
}
