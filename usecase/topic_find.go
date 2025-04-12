package usecase

import (
	"context"
	"englishAI/entities"
	"englishAI/repository"
	"fmt"
)

type topicFindUseCase struct {
	topicRepo repository.ITopicRepository
}

func NewTopicFindUseCase(topicRepo repository.ITopicRepository) ITopicFindUsecase {
	return &topicFindUseCase{topicRepo: topicRepo}
}

func (uc *topicFindUseCase) Execute(ctx context.Context, paging entities.PagingRequest) ([]entities.Topic, uint32, error) {
	total, err := uc.topicRepo.CountTotal(ctx)
	if err != nil {
		return nil, 0, fmt.Errorf("can not find topics: %w", err)
	}

	if total == 0 {
		return []entities.Topic{}, 0, nil
	}

	res, err := uc.topicRepo.GetByPage(ctx, paging)
	if err != nil {
		return nil, 0, fmt.Errorf("can not get topics: %w", err)
	}

	if len(res) == 0 {
		return []entities.Topic{}, 0, nil
	}

	return res, total, nil
}
