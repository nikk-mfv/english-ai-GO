package usecase

import (
	"englishAI/entities"
	"englishAI/repository"

	"github.com/gin-gonic/gin"
)

type topicCreateUsecase struct {
	topicRepository repository.ITopicRepository
}

func NewTopicCreateUsecase(topicRepository repository.ITopicRepository) ITopicCreateUsecase {
	return &topicCreateUsecase{topicRepository: topicRepository}
}

func (uc *topicCreateUsecase) Execute(ctx *gin.Context, topic entities.Topic) (entities.Topic, error) {

	if err := uc.topicRepository.CreateTopic(ctx, &topic); err != nil {
		return entities.Topic{}, err
	}
	return topic, nil
}
