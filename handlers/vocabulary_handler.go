package handlers

import (
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

	vocab, err := ucCreate.Execute(ctx, input.Name, input.Definition, input.Example, input.Pronunciation)

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
