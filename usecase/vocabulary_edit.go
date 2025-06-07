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

func (uc *vocabularyUpdateUsecase) Execute(ctx context.Context, id string, name string, definition string, example string, pronunciation string, topicIds []uint) (entities.Vocabulary, error) {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	var vocab, err = uc.vocabularyRepo.UpdateByID(ctx, id, name, definition, example, pronunciation, topicIds)
	if err != nil {
		return entities.Vocabulary{}, err
	}

	return vocab, nil
}
