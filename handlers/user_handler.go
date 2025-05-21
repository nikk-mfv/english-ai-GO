package handlers

import (
	"englishAI/entities"
	"englishAI/repository"
	"englishAI/usecase"
	"englishAI/utils"
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

func (h *userHandler) SignUp(ctx *gin.Context) {
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

func (h *userHandler) Login(ctx *gin.Context) {
	var input struct {
		UserName string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := ucUserFind.Execute(ctx, input.UserName)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error finding user"})
		return
	} else if user == nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
		return
	}

	token, err := utils.GenerateToken(user.ID, user.Username)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error generating token"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Login successful", "token": token})
}

func (h *userHandler) Profile(ctx *gin.Context) {
	userID := ctx.MustGet("user_id").(uint)
	username := ctx.MustGet("username").(string)

	ctx.JSON(http.StatusOK, gin.H{
		"user_id":  userID,
		"username": username,
	})
}
