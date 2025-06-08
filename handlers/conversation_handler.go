package handlers

import (
	"englishAI/entities"
	"englishAI/repository"
	"englishAI/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type conversationHandler struct {
	BaseHandler
}

func NewConversationHandler() *conversationHandler {
	return &conversationHandler{}
}

// conversation repo
var conversationRepo = repository.NewConversationRepository()

// conversation usecase
var createUsecase = usecase.NewConversationCreateUsecase(conversationRepo)
var findUsecase = usecase.NewConversationFindUseCase(conversationRepo)
var findAllUsecase = usecase.NewConversationFindAllUsecase(conversationRepo)

func (h *conversationHandler) Find(ctx *gin.Context) {
	id := ctx.Param("id")

	var conversation = entities.Conversation{}

	// call usecase
	conversation, err := findUsecase.Execute(id)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, entities.Response{
		Data: conversation,
	})

}

func (h *conversationHandler) Create(c *gin.Context) {
	var input struct {
		Name string `json:"name" binding:"required"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	UserIDLogged := c.GetUint("user_id")
	conversation := entities.Conversation{
		Name:   input.Name,
		UserID: UserIDLogged,
	}

	conversation, err := createUsecase.Execute(c, conversation)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": conversation})
}

func (h *conversationHandler) FindAll(ctx *gin.Context) {
	UserIDLogged := ctx.GetUint("user_id")
	var conversations = []entities.Conversation{}

	// call usecase
	conversations, err := findAllUsecase.Execute(UserIDLogged)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, entities.Response{
		Data: conversations,
	})
}
