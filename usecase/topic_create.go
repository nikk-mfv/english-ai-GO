package usecase

import (
	"context"
	"englishAI/entities"
	"englishAI/repository"
)

type topicCreateUsecase struct {
	topicRepository repository.ITopicRepository
}

func NewTopicCreateUsecase(topicRepository repository.ITopicRepository) ITopicCreateUsecase {
	return &topicCreateUsecase{topicRepository: topicRepository}
}

func (uc *topicCreateUsecase) Execute(ctx context.Context, topic entities.Topic) (entities.Topic, error) {
	if err := uc.topicRepository.Create(ctx, &topic); err != nil {
		return entities.Topic{}, err
	}
	return topic, nil
}
