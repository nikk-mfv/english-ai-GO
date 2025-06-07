package repository

import (
	"context"
	"englishAI/config"
	"englishAI/entities"
)

type topicRepository struct{}

func NewTopicRepository() ITopicRepository {
	return &topicRepository{}
}

func (r *topicRepository) CountTotal(ctx context.Context) (uint32, error) {
	var totalTopics int64
	db := config.GetDatabase()

	if err := db.Model(&entities.Topic{}).Count(&totalTopics).Error; err != nil {
		return 0, err
	}

	return uint32(totalTopics), nil
}

func (r *topicRepository) Create(ctx context.Context, newtopic *entities.Topic) error {
	return config.GetDatabase().Create(newtopic).Error
}

func (r *topicRepository) GetByPage(ctx context.Context, paging entities.PagingRequest) ([]entities.Topic, error) {
	db := config.GetDatabase()

	var topics []entities.Topic
	gormPaging := paging.GormPaging(db)
	if err := gormPaging.Find(&topics).Error; err != nil {
		return nil, err
	}

	return topics, nil
}

func (r *topicRepository) UpdateByID(ctx context.Context, id string, name string) (entities.Topic, error) {
	var topic entities.Topic
	db := config.GetDatabase()

	if err := db.Model(&topic).Where("id = ?", id).Updates(entities.Topic{Name: name}).Error; err != nil {
		return topic, err
	}

	if err := db.Where("id = ?", id).First(&topic).Error; err != nil {
		return topic, err
	}

	return topic, nil
}

func (r *topicRepository) DeleteByID(ctx context.Context, id string) error {
	db := config.GetDatabase()
	var topic entities.Topic

	if err := db.Where("id = ?", id).First(&topic).Error; err != nil {
		return err
	}

	if err := db.Unscoped().Where("id = ?", id).Delete(&topic).Error; err != nil {
		return err
	}

	return nil
}
