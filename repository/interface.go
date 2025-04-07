package repository

import (
	"context"
	"englishAI/entities"

	"github.com/gin-gonic/gin"
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
	Create(ctx *gin.Context, topic *entities.Topic) error
}
