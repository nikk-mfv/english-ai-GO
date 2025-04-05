package usecase

import (
	"context"
	"englishAI/entities"
	"englishAI/repository"
	"time"
)

type vocabularyUpdateUsecase struct {
	vocabularyRepo repository.IVocabularyRepository
}

func NewVocabularyUpdateUsecase(vocabularyRepo repository.IVocabularyRepository) IVocabularyUpdateUsecase {
	return &vocabularyUpdateUsecase{
		vocabularyRepo: vocabularyRepo,
	}
}

func (uc *vocabularyUpdateUsecase) Execute(ctx context.Context, id string, obj *entities.Vocabulary) error {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	if err := uc.vocabularyRepo.UpdateByID(ctx, id, obj); err != nil {
		return err
	}

	return uc.vocabularyRepo.UpdateByID(ctx, id, obj)
}
