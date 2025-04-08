package usecase

import (
	"englishAI/entities"
	"englishAI/repository"

	"github.com/gin-gonic/gin"
)

type topicFindUseCase struct {
	topicRepo repository.ITopicRepository
}

func NewTopicFindUseCase(topicRepo repository.ITopicRepository) ITopicFindUsecase {
	return &topicFindUseCase{topicRepo: topicRepo}
}

func (uc *topicFindUseCase) Execute(ctx *gin.Context) ([]entities.Topic, error) {
	return uc.topicRepo.GetAll(ctx)
}
