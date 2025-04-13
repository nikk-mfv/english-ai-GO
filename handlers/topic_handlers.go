package handlers

import (
	"englishAI/entities"
	"englishAI/repository"
	"englishAI/usecase"
	"net/http"
	"strconv"

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
	pageStr := ctx.DefaultQuery("page", "1")
	sizeStr := ctx.DefaultQuery("size", "10")

	page, err := strconv.ParseInt(pageStr, 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid page param"})
		return
	}

	size, err := strconv.ParseInt(sizeStr, 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid size param"})
		return
	}

	paging := entities.PagingRequest{
		Page: uint32(page),
		Size: uint32(size),
	}

	topics, total, err := ucFindTopicsByPage.Execute(ctx, paging)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "cannot find topics: " + err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, entities.Response{
		Data:   topics,
		Paging: entities.ResponsePaging{Total: total},
	})
}
