package repository

import (
	"context"
	"englishAI/config"
	"englishAI/entities"
)

type vocabularyRepository struct{}

func NewVocabularyRepository() IVocabularyRepository {
	return &vocabularyRepository{}
}

func (r *vocabularyRepository) Count(ctx context.Context) (uint32, error) {
	var count int64
	if err := config.GetDatabase().Model(&entities.Vocabulary{}).Count(&count).Error; err != nil {
		return 0, err
	}

	return uint32(count), nil
}

func (r *vocabularyRepository) Create(ctx context.Context, name string, definition string, example string, pronunciation string, topicIds []uint) (entities.Vocabulary, error) {
	var db = config.GetDatabase()

	var topics []*entities.Topic
	if err := config.GetDatabase().Where("id IN ?", topicIds).Find(&topics).Error; err != nil {
		return entities.Vocabulary{}, err
	}

	var newVocab = entities.Vocabulary{
		Name:          name,
		Definition:    definition,
		Example:       example,
		Pronunciation: pronunciation,
		Topics:        topics,
	}

	if err := db.Create(&newVocab).Error; err != nil {
		return entities.Vocabulary{}, err
	}

	return newVocab, nil
}

func (r *vocabularyRepository) GetAll(ctx context.Context, paging entities.PagingRequest) ([]entities.Vocabulary, error) {
	var vocabularies []entities.Vocabulary

	GormPaging := paging.GormPaging(config.GetDatabase())

	if err := GormPaging.Preload("Topics").Find(&vocabularies).Error; err != nil {
		return nil, err
	}
	return vocabularies, nil
}

func (r *vocabularyRepository) DeleteByID(ctx context.Context, id string) error {
	if err := config.GetDatabase().First(&entities.Vocabulary{}, "id = ?", id).Error; err != nil {
		return err
	}

	return config.GetDatabase().Unscoped().Where("id = ?", id).Delete(&entities.Vocabulary{}).Error
}

func (r *vocabularyRepository) UpdateByID(ctx context.Context, id string, name string, definition string, example string, pronunciation string, topicIds []uint) (entities.Vocabulary, error) {
	var vocab entities.Vocabulary
	if err := config.GetDatabase().Preload("Topics").First(&vocab, "id = ?", id).Error; err != nil {
		return entities.Vocabulary{}, err
	}

	var topics []*entities.Topic
	if err := config.GetDatabase().Where("id IN ?", topicIds).Find(&topics).Error; err != nil {
		return entities.Vocabulary{}, err
	}

	vocab.Name = name
	vocab.Definition = definition
	vocab.Example = example
	vocab.Pronunciation = pronunciation

	if err := config.GetDatabase().Model(&vocab).Association("Topics").Replace(topics); err != nil {
		return entities.Vocabulary{}, err
	}

	if err := config.GetDatabase().Save(&vocab).Error; err != nil {
		return entities.Vocabulary{}, err
	}

	return vocab, nil
}
