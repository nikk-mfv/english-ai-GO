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

func (uc *vocabularyCreateUsecase) Execute(ctx context.Context, Name string, Definition string, Example string, Pronunciation string) ([]entities.Vocabulary, error) {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	newVocabulary := entities.Vocabulary{
		Name:          Name,
		Definition:    Definition,
		Example:       Example,
		Pronunciation: Pronunciation,
	}

	if err := uc.vocabularyRepo.Create(ctx, &newVocabulary); err != nil {
		return nil, err
	}

	return []entities.Vocabulary{newVocabulary}, nil
}
