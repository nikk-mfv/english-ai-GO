package repository

import (
	"context"
	"englishAI/entities"
)

type IVocabularyRepository interface {
	Create(ctx context.Context, obj *entities.Vocabulary) error
	GetAll(ctx context.Context) ([]entities.Vocabulary, error)
}
