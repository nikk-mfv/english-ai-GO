package usecase

import (
	"context"
	"englishAI/entities"
	"englishAI/repository"
	"fmt"
)

type vocabularyFindUsecase struct {
	vocabularyRepo repository.IVocabularyRepository
}

func NewVocabularyFindUsecase(vocabularyRepo repository.IVocabularyRepository) IVocabularyFindUsecase {
	return &vocabularyFindUsecase{vocabularyRepo: vocabularyRepo}
}

func (uc *vocabularyFindUsecase) Execute(ctx context.Context, paging entities.PagingRequest) ([]entities.Vocabulary, uint32, error) {
	total, err := uc.vocabularyRepo.Count(ctx)
	if err != nil {
		return nil, 0, fmt.Errorf("can not find vocab: %w", err)
	}

	if total == 0 {
		return nil, 0, nil
	}

	res, err := uc.vocabularyRepo.GetAll(ctx, paging)

	if err != nil {
		return nil, 0, fmt.Errorf("can not get vocab: %w", err)
	}

	if len(res) == 0 {
		return nil, 0, nil
	}

	return res, total, err
}
