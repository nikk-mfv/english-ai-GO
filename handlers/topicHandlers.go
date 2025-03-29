package handlers

import (
	"englishAI/repository"

	"github.com/gin-gonic/gin"
)

type TopicHandler struct{}

var topicRepository = repository.TopicRepository{}

func (h *TopicHandler) GetTopics(c *gin.Context) {
	topicRepository.GetTopics(c)
}

func (h *TopicHandler) CreateTopics(c *gin.Context) {
	topicRepository.CreateTopics(c)
}
