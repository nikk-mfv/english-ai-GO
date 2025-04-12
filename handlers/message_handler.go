package handlers

import (
	"englishAI/entities"
	"englishAI/external"
	"englishAI/repository"
	"englishAI/usecase"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type messageHandler struct {
	BaseHandler
}

func NewMessageHandler() *messageHandler {
	return &messageHandler{}
}

var messageRepo = repository.NewMessageRepo()
var aiService = external.NewAiService()
var createMessageUsecase = usecase.NewMessageCreateUsecase(messageRepo)

func (h *messageHandler) Create(ctx *gin.Context) {
	var input struct {
		Message        string `json:"message" binding:"required"`
		ConversationID uint   `json:"conversation_id" binding:"required"`
	}

	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var newMess = &entities.Message{
		Message:        input.Message,
		ConversationID: input.ConversationID,
		IsHuman:        true,
		UserID:         1, // TODO: fix it after handling AUTHENTICATE
	}

	log.Printf(newMess.Message)

	// call usecase to store message from human
	err := createMessageUsecase.Execute(newMess)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// get res from AI serivce
	aiMessStr, errAI := aiService.Reply(newMess.Message)
	if errAI != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "error from AI"})
		return
	}

	log.Println(aiMessStr)
	var aiMess = &entities.Message{
		Message:        aiMessStr,
		ConversationID: input.ConversationID,
		IsHuman:        false,
		UserID:         1, // TODO: fix it after handling AUTHENTICATE
	}

	err = createMessageUsecase.Execute(aiMess)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// return message from AI
	ctx.JSON(http.StatusCreated, entities.Response{
		Data: aiMess,
	})
}
