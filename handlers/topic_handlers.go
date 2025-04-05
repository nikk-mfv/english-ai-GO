package handlers

import (
	"englishAI/config"
	"englishAI/entities"
	"englishAI/repository"
	"englishAI/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type topicHandler struct {
	BaseHandler
}

var (
	// repo
	topicRepository = repository.NewTopicRepository()
	// usecase
	topicCreateUsecase = usecase.NewTopicCreateUsecase(topicRepository)
	topicGetUsecase    = usecase.NewTopicGetAllUsecase(topicRepository)
)

func NewTopicHandler() *topicHandler {
	return &topicHandler{}
}

func (hd1 *topicHandler) Create(ctx *gin.Context) {

	var topicInput struct {
		Name   string `json:"name" binding:"required"`
		UserID uint   `json:"user_id" binding:"required"`
	}

	// Bind JSON to newtopic
	if err := ctx.ShouldBindJSON(&topicInput); err != nil {
		respondWithJSON(ctx, http.StatusBadRequest, gin.H{"error": "invalid request data"})
		return
	}

	newTopic := entities.Topic{
		Name:   topicInput.Name,
		UserID: topicInput.UserID,
	}

	// check existing topic
	var existingTopic entities.Topic
	if err := config.GetDatabase().Where("name= ? ", newTopic.Name).First(&existingTopic).Error; err == nil {
		respondWithJSON(ctx, http.StatusBadRequest, gin.H{"error": "topic already existed"})
		return
	}

	createdTopic, err := topicCreateUsecase.Execute(ctx, newTopic)
	//add new topic to database
	if err != nil {
		respondWithJSON(ctx, http.StatusInternalServerError, gin.H{"error": "cannot create new topic: " + err.Error()})
		return
	}

	//success, do not have error
	respondWithJSON(ctx, http.StatusOK, createdTopic)
}

func (hd1 *topicHandler) Find(ctx *gin.Context) {
	topics, err := topicGetUsecase.Execute(ctx)

	if err != nil {
		respondWithJSON(ctx, http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	//success, do not have error
	respondWithJSON(ctx, http.StatusOK, topics)
}
