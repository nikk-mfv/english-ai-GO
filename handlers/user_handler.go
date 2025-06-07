package handlers

import (
	"englishAI/entities"
	"englishAI/repository"
	"englishAI/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type userHandler struct {
	BaseHandler
}

var userRepo = repository.NewUserRepository()

var ucUserCreate = usecase.NewUserCreateUsecase(userRepo)
var ucUserFind = usecase.NewUserFindUsecase(userRepo)

func NewUserHandler() *userHandler {
	return &userHandler{}
}

func (h *userHandler) Create(ctx *gin.Context) {
	var input struct {
		UserName string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required,min=6"`
	}

	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if existingUser, err := ucUserFind.Execute(ctx, input.UserName); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error checking existing user"})
		return
	} else if existingUser != nil {
		ctx.JSON(http.StatusConflict, gin.H{"error": "Username already exists"})
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error hashing password"})
		return
	}

	user := entities.User{
		Username: input.UserName,
		Password: string(hashedPassword),
	}

	if err := ucUserCreate.Execute(ctx, user); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "User created successfully"})
}
