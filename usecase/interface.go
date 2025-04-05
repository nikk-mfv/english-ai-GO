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

type IVocabularyDeleteUsecase interface {
	Execute(ctx context.Context, id string) error
}

type IVocabularyUpdateUsecase interface {
	Execute(ctx context.Context, id string, obj *entities.Vocabulary) error
}
