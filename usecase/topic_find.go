package usecase

import (
	"context"
	"englishAI/entities"
	"englishAI/repository"
)

type topicFindUseCase struct {
	topicRepo repository.ITopicRepository
}

func NewTopicFindUseCase(topicRepo repository.ITopicRepository) ITopicFindUsecase {
	return &topicFindUseCase{topicRepo: topicRepo}
}

func (uc *topicFindUseCase) Execute(ctx context.Context, paging entities.PagingRequest) ([]entities.Topic, int64, error) {
	return uc.topicRepo.GetByPage(ctx, paging)
}
