package usecase

import (
	"context"
	"englishAI/entities"
	"englishAI/repository"
	"time"
)

type vocabularyCreateUsecase struct {
	vocabularyRepo repository.IVocabularyRepository
}

func NewVocabularyCreateUsecase(
	vocabularyRepo repository.IVocabularyRepository,
) IVocabularyCreateUsecase {
	return &vocabularyCreateUsecase{
		vocabularyRepo: vocabularyRepo,
	}
}

func (uc *vocabularyCreateUsecase) Execute(ctx context.Context, newVocabulary entities.Vocabulary) (entities.Vocabulary, error) {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	if err := uc.vocabularyRepo.Create(ctx, &newVocabulary); err != nil {
		return entities.Vocabulary{}, err
	}

	return newVocabulary, nil
}
