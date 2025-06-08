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

type VocabularyHandler struct {
	BaseHandler
}

// repo
var vocabRepo = repository.NewVocabularyRepository()

// usecase
var ucCreate = usecase.NewVocabularyCreateUsecase(vocabRepo)
var ucFind = usecase.NewVocabularyFindUsecase(vocabRepo)
var ucDelete = usecase.NewVocabularyDeleteUsecase(vocabRepo)
var ucUpdate = usecase.NewVocabularyUpdateUsecase(vocabRepo)

func NewVocabularyHandler() *VocabularyHandler {
	return &VocabularyHandler{}
}

func (hdl *VocabularyHandler) Create(ctx *gin.Context) {
	var input struct {
		Name          string `json:"name" binding:"required"`
		Definition    string `json:"definition" binding:"required"`
		Example       string `json:"example"`
		Pronunciation string `json:"pronunciation"`
		TopicIds      []uint `json:"topic_ids"`
	}

	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID := ctx.GetUint("user_id")
	vocab, err := ucCreate.Execute(ctx, input.Name,
		input.Definition,
		input.Example,
		input.Pronunciation,
		input.TopicIds,
		userID)

	if err != nil {
		if strings.Contains(err.Error(), "Duplicate entry") {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Vocabulary already exists"})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"vocabulary": vocab, "message": "Vocabulary created successfully"})
}

func (hdl *VocabularyHandler) Find(ctx *gin.Context) {
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
	vocab, total, err := ucFind.Execute(ctx, paging, userID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, entities.Response{
		Data:   vocab,
		Paging: entities.ResponsePaging{Total: total},
	})
}

func (hdl *VocabularyHandler) DeleteById(ctx *gin.Context) {
	id := ctx.Param("id")

	if err := ucDelete.Execute(ctx, id); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Vocabulary deleted successfully"})
}

func (hdl *VocabularyHandler) UpdateById(ctx *gin.Context) {
	id := ctx.Param("id")

	var input struct {
		Name          string `json:"name"`
		Definition    string `json:"definition"`
		Example       string `json:"example"`
		Pronunciation string `json:"pronunciation"`
		TopicIds      []uint `json:"topic_ids"`
	}

	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	vocab, err := ucUpdate.Execute(ctx, id, input.Name, input.Definition, input.Example, input.Pronunciation, input.TopicIds)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"vocabulary": vocab, "message": "Vocabulary updated successfully"})
}
