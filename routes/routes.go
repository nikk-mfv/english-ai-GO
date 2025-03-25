package routes

import (
	"englishAI/handlers"

	"github.com/gin-gonic/gin"
)

func Routes(r *gin.Engine) {

	topicHandler := handlers.TopicHandler{}
	TopicGroup := r.Group("/api/v1/topics")
	{
		TopicGroup.GET("/", topicHandler.GetTopics)
		TopicGroup.POST("/", topicHandler.CreateTopics)
	}
}
