package handlers

import (
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
	ucCreateTopic      = usecase.NewTopicCreateUsecase(topicRepository)
	ucFindTopicsByPage = usecase.NewTopicFindUseCase(topicRepository)
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
	var paging entities.PagingRequest

	if err := ctx.ShouldBindQuery(&paging); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid params"})
		return
	}

	paging.SetDefautls()

	topics, total, err := ucFindTopicsByPage.Execute(ctx, paging)
	if err != nil {
		if err.Error() == "topic already exists" {
			ctx.JSON(http.StatusConflict, gin.H{"error": err.Error()})
		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "cannot find topics: " + err.Error()})
		}
		return
	}

	ctx.JSON(http.StatusOK, entities.PagingResponse[entities.Topic]{
		Data:       topics,
		TotalItems: total,
		Size:       int64(paging.Size),
	})
}
