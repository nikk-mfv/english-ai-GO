package repository

import (
	"context"
	"englishAI/entities"

	"github.com/gin-gonic/gin"
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
	GetTopics(ctx *gin.Context) ([]entities.Topic, error)
	CreateTopic(ctx *gin.Context, topic *entities.Topic) error
}
