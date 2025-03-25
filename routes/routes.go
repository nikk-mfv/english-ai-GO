package routes

import (
	"englishAI/handlers"

	"github.com/gin-gonic/gin"
)

func Routes(r *gin.Engine) {
	userGroup := r.Group("/api/v1/users")
	{
		userGroup.GET("/", handlers.GetUsers)
		userGroup.POST("/", handlers.CreateUser)
	}
	TopicGroup := r.Group("/api/v1/topics")
	{
		TopicGroup.GET("/", handlers.GetTopics)
	}
}
