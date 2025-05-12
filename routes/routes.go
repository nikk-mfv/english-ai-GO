package routes

import (
	"englishAI/handlers"

	"github.com/gin-gonic/gin"
)

func Routes(r *gin.Engine) {
	topicHandler := handlers.NewTopicHandler()
	TopicGroup := r.Group("/api/v1/topic")
	{
		TopicGroup.POST("", topicHandler.Create)
		TopicGroup.GET("", topicHandler.Find)
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
