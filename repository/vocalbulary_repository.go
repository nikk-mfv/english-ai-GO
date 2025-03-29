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

func (r *vocabularyRepository) Create(ctx context.Context, obj *entities.Vocabulary) error {
	return config.GetDatabase().Create(obj).Error
}

func (r *vocabularyRepository) GetAll(ctx context.Context) ([]entities.Vocabulary, error) {
	var vocabularies []entities.Vocabulary
	if err := config.GetDatabase().Find(&vocabularies).Error; err != nil {
		return nil, err
	}
	return vocabularies, nil
}
