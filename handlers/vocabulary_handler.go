package handlers

import (
	"englishAI/entities"
	"englishAI/repository"
	"englishAI/usecase"
	"net/http"

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
	}

	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	vocab, err := ucCreate.Execute(ctx, entities.Vocabulary{
		Name:          input.Name,
		Definition:    input.Definition,
		Example:       input.Example,
		Pronunciation: input.Pronunciation,
	})

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, vocab)
}

func (hdl *VocabularyHandler) Find(ctx *gin.Context) {
	vocab, err := ucFind.Execute(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, vocab)
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
	}

	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	vocab := entities.Vocabulary{
		Name:          input.Name,
		Definition:    input.Definition,
		Example:       input.Example,
		Pronunciation: input.Pronunciation,
	}

	if err := ucUpdate.Execute(ctx, id, &vocab); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Vocabulary updated successfully"})
}
