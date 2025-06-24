package routes

import (
	"englishAI/handlers"

	"englishAI/middleware"

	"github.com/gin-gonic/gin"
)

func Routes(r *gin.Engine) {
	// auth
	userHandler := handlers.NewUserHandler()
	UserGroup := r.Group("/api/v1/user")
	{
		UserGroup.POST("/create-account", userHandler.SignUp)
		UserGroup.POST("/log-in", userHandler.Login)
		UserGroup.GET("/profile", middleware.AuthMiddleware(), userHandler.Profile)
		UserGroup.POST("/upload-avatar", middleware.AuthMiddleware(), userHandler.UploadAvatar)
	}
	// apply middleware
	r.Use(middleware.AuthMiddleware())

	// routers ...

	// Topic
	topicHandler := handlers.NewTopicHandler()
	TopicGroup := r.Group("/api/v1/topic")
	{
		TopicGroup.POST("", topicHandler.Create)
		TopicGroup.GET("", topicHandler.Find)
		TopicGroup.PUT("/:id", topicHandler.UpdateById)
		TopicGroup.DELETE("/:id", topicHandler.DeleteById)
	}

	vocabHandler := handlers.VocabularyHandler{}
	VocabGroup := r.Group("/api/v1/vocab")
	{
		VocabGroup.POST("", vocabHandler.Create)
		VocabGroup.GET("", vocabHandler.Find)
		VocabGroup.DELETE("/:id", vocabHandler.DeleteById)
		VocabGroup.PUT("/:id", vocabHandler.UpdateById)
	}

	// CONVERSATION
	conversationHandler := handlers.NewConversationHandler()
	ConversationGroup := r.Group("/api/v1/conversation")
	{
		ConversationGroup.POST("", conversationHandler.Create)
		ConversationGroup.GET("/:id", conversationHandler.Find)
		ConversationGroup.GET("", conversationHandler.FindAll)
	}

	// messages
	messageHandler := handlers.NewMessageHandler()
	MessageGroup := r.Group("/api/v1/message")
	{
		MessageGroup.POST("", messageHandler.Create)
		MessageGroup.GET("", messageHandler.FindAll)
	}

}
