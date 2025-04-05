package usecase

import (
	"englishAI/entities"
	"englishAI/repository"

	"github.com/gin-gonic/gin"
)

type topicGetAllUsecase struct {
	topicRepository repository.ITopicRepository
}

func NewTopicGetAllUsecase(topicRepository repository.ITopicRepository) ITopicGetAllUsecase {
	return &topicGetAllUsecase{topicRepository: topicRepository}
}

func (uc *topicGetAllUsecase) Execute(ctx *gin.Context) ([]entities.Topic, error) {

	topics, err := uc.topicRepository.GetTopics(ctx)
	if err != nil {
		return nil, err
	}
	return topics, nil
}
