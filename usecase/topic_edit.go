package usecase

import (
	"context"
	"englishAI/entities"
	"englishAI/repository"
)

type topicUpdateUsecase struct {
	topicRepo repository.ITopicRepository
}

func NewTopicUpdateUsecase(topicRepo repository.ITopicRepository) ITopicUpdateUsecase {
	return &topicUpdateUsecase{
		topicRepo: topicRepo,
	}
}

func (uc *topicUpdateUsecase) Execute(ctx context.Context, id string, name string) (entities.Topic, error) {
	topic, err := uc.topicRepo.UpdateByID(ctx, id, name)
	if err != nil {
		return entities.Topic{}, err
	}

	return topic, nil
}
