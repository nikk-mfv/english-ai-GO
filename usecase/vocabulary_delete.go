package usecase

import (
	"context"
	"englishAI/repository"
	"time"
)

type vocabularyDeleteUsecase struct {
	vocabularyRepo repository.IVocabularyRepository
}

func NewVocabularyDeleteUsecase(vocabularyRepo repository.IVocabularyRepository) IVocabularyDeleteUsecase {
	return &vocabularyDeleteUsecase{
		vocabularyRepo: vocabularyRepo,
	}
}

func (uc *vocabularyDeleteUsecase) Execute(ctx context.Context, id string) error {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	if err := uc.vocabularyRepo.DeleteByID(ctx, id); err != nil {
		return err
	}

	return nil
}
