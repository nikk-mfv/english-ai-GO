package usecase

import (
	"context"
	"englishAI/repository"
)

type TopicDeleteUsecase struct {
	topicRepo repository.ITopicRepository
}

func NewTopicDeleteUsecase(topicRepo repository.ITopicRepository) ITopicDeleteUsecase {
	return &TopicDeleteUsecase{
		topicRepo: topicRepo,
	}
}

func (uc *TopicDeleteUsecase) Execute(ctx context.Context, id string) error {
	if err := uc.topicRepo.DeleteByID(ctx, id); err != nil {
		return err
	}
	return nil
}
