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

func (uc *vocabularyCreateUsecase) Execute(ctx context.Context, name string, definition string, example string, pronunciation string, topicIds []uint, userID uint) (entities.Vocabulary, error) {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	vocab, err := uc.vocabularyRepo.Create(ctx, name, definition, example, pronunciation, topicIds, userID)

	if err != nil {
		return entities.Vocabulary{}, err
	}
	return vocab, nil
}
