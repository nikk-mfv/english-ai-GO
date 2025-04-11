package repository

import (
	"context"
	"englishAI/config"
	"englishAI/entities"
	"errors"

	"gorm.io/gorm"
)

type topicRepository struct{}

func NewTopicRepository() ITopicRepository {
	return &topicRepository{}
}

func (r *topicRepository) Create(ctx context.Context, newtopic *entities.Topic) error {
	if err := r.CheckExisted(ctx, newtopic); err != nil {
		return err
	}

	return config.GetDatabase().Create(newtopic).Error
}

func (r *topicRepository) GetByPage(ctx context.Context, paging entities.PagingRequest) ([]entities.Topic, int64, error) {
	db := config.GetDatabase()
	var topics []entities.Topic
	var totalTopics int64
	offset := (paging.Page - 1) * paging.Size

	if err := db.Model(&entities.Topic{}).Count(&totalTopics).Error; err != nil {
		return nil, 0, err
	}

	if err := db.Order("created_at DESC").Limit(int(paging.Size)).Offset(int(offset)).Find(&topics).Error; err != nil {
		return nil, 0, err
	}

	return topics, totalTopics, nil
}

func (r *topicRepository) CheckExisted(ctx context.Context, newtopic *entities.Topic) error {
	db := config.GetDatabase()

	var ErrTopicAlreadyExists = errors.New("topic already exists")
	var existingTopic entities.Topic
	err := db.Where("name= ? ", newtopic.Name).First(&existingTopic).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil
		}
		return err
	}
	return ErrTopicAlreadyExists
}
