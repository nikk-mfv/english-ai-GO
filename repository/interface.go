package repository

import (
	"context"
	"englishAI/entities"
)

type IVocabularyRepository interface {
	Count(ctx context.Context) (uint32, error)
	Create(ctx context.Context, obj *entities.Vocabulary) error
	GetAll(ctx context.Context, paging entities.PagingRequest) ([]entities.Vocabulary, error)
	DeleteByID(ctx context.Context, id string) error
	UpdateByID(ctx context.Context, id string, obj *entities.Vocabulary) error
}

type IConversationRepository interface {
	Create(ctx context.Context, obj *entities.Conversation) error
}

type ITopicRepository interface {
	CountTotal(ctx context.Context) (uint32, error)
	Create(ctx context.Context, topic *entities.Topic) error
	GetByPage(ctx context.Context, paging entities.PagingRequest) ([]entities.Topic, error)
}
