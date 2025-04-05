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

func (r *topicRepository) GetTopics(ctx *gin.Context) ([]entities.Topic, error) {
	var topics []entities.Topic
	if err := config.GetDatabase().Find(&topics).Error; err != nil {
		return nil, err
	}
	return topics, nil
}

func (r *topicRepository) CreateTopic(ctx *gin.Context, newtopic *entities.Topic) error {
	return config.GetDatabase().Create(newtopic).Error
}
