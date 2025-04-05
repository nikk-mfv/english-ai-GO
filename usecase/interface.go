package usecase

import (
	"context"
	"englishAI/entities"
)

type IVocabularyCreateUsecase interface {
	Execute(ctx context.Context, Name string, Definition string, Example string, Pronunciation string) ([]entities.Vocabulary, error)
}

type IVocabularyFindUsecase interface {
	Execute(ctx context.Context) ([]entities.Vocabulary, error)
}
