package repository

import (
	"englishAI/config"
	"englishAI/entities"

	"github.com/gin-gonic/gin"
)

type topicRepository struct{}

func NewTopicRepository() ITopicRepository {
	return &topicRepository{}
}

func (r *topicRepository) Create(ctx *gin.Context, newtopic *entities.Topic) error {
	return config.GetDatabase().Create(newtopic).Error
}
