package usecase

import (
	"context"
	"englishAI/entities"
	"englishAI/repository"
)

type vocabularyFindUsecase struct {
	vocabularyRepo repository.IVocabularyRepository
}

func NewVocabularyFindUsecase(vocabularyRepo repository.IVocabularyRepository) IVocabularyFindUsecase {
	return &vocabularyFindUsecase{vocabularyRepo: vocabularyRepo}
}

func (uc *vocabularyFindUsecase) Execute(ctx context.Context) ([]entities.Vocabulary, error) {
	return uc.vocabularyRepo.GetAll(ctx)
}
