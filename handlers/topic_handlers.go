package handlers

import (
	"englishAI/entities"
	"englishAI/repository"
	"englishAI/usecase"
	"net/http"
	"strconv"
	"strings"

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
	ucUpdateTopic      = usecase.NewTopicUpdateUsecase(topicRepository)
	ucDeleteTopic      = usecase.NewTopicDeleteUsecase(topicRepository)
)

func NewTopicHandler() *topicHandler {
	return &topicHandler{}
}

func (hd1 *topicHandler) Create(ctx *gin.Context) {
	var topicInput struct {
		Name string `json:"name" binding:"required"`
	}

	// Bind JSON to newtopic
	if err := ctx.ShouldBindJSON(&topicInput); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid request data" + err.Error()})
		return
	}

	userID := ctx.GetUint("user_id")

	newTopic := entities.Topic{
		Name:   topicInput.Name,
		UserID: userID,
	}

	//add new topic to database
	createdTopic, err := ucCreateTopic.Execute(ctx, newTopic)
	if err != nil {
		if strings.Contains(err.Error(), "Duplicate entry") {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "topic already exists"})
			return
		}
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

	userID := ctx.GetUint("user_id")

	topics, total, err := ucFindTopicsByPage.Execute(ctx, paging, userID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "cannot find topics: " + err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, entities.Response{
		Data:   topics,
		Paging: entities.ResponsePaging{Total: total},
	})
}

func (hd1 *topicHandler) UpdateById(ctx *gin.Context) {
	topicId := ctx.Param("id")
	if topicId == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid topic id"})
		return
	}

	var topicInput struct {
		Name string `json:"name" binding:"required"`
	}

	if err := ctx.ShouldBindJSON(&topicInput); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid request data" + err.Error()})
		return
	}

	updatedTopic, err := ucUpdateTopic.Execute(ctx, topicId, topicInput.Name)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "cannot update topic: " + err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"topic": updatedTopic, "message": "topic updated successfully"})
}

func (hd1 *topicHandler) DeleteById(ctx *gin.Context) {
	topicId := ctx.Param("id")

	if topicId == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid topic id"})
		return
	}

	err := ucDeleteTopic.Execute(ctx, topicId)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "cannot delete topic: " + err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "topic deleted successfully"})
}
