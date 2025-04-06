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
	ucCreateTopic = usecase.NewTopicCreateUsecase(topicRepository)
	ucFindTopics  = usecase.NewTopicFindUseCase(topicRepository)
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
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid request data" + err.Error()})
		return
	}

	newTopic := entities.Topic{
		Name:   topicInput.Name,
		UserID: topicInput.UserID,
	}

	// check existing topic
	var existingTopic entities.Topic
	if err := config.GetDatabase().Where("name= ? ", newTopic.Name).First(&existingTopic).Error; err == nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "topic already existed"})
		return
	}

	//add new topic to database
	createdTopic, err := ucCreateTopic.Execute(ctx, newTopic)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "cannot create new topic: " + err.Error()})
		return
	}

	//successfully created topic
	ctx.JSON(http.StatusOK, createdTopic)
}

func (hd1 *topicHandler) Find(ctx *gin.Context) {
	topics, err := ucFindTopics.Execute(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "cannot find topics: " + err.Error()})
		return
	}

	if len(topics) == 0 {
		ctx.JSON(http.StatusNotFound, gin.H{"message": "no topics found"})
		return
	}

	ctx.JSON(http.StatusOK, topics)
}
