package handlers

import (
	"englishAI/entities"
	"englishAI/repository"
	"englishAI/usecase"
	"englishAI/utils"
	"fmt"
	"net/http"
	"regexp"
	"strings"

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

type UserInput struct {
	UserName string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func validateUserInput(ctx *gin.Context) (UserInput, error) {
	var input UserInput

	if err := ctx.ShouldBindJSON(&input); err != nil {
		return input, fmt.Errorf("invalid input: %v", err)
	}

	if len(input.UserName) < 3 || len(input.UserName) > 30 {
		return input, fmt.Errorf("name must be between 3 and 30 characters")
	}

	usernameRegex := regexp.MustCompile(`^[A-Za-z][A-Za-z0-9\-]*$`)
	if !usernameRegex.MatchString(input.UserName) {
		return input, fmt.Errorf("name can only contain letters, numbers, or dashes, and must start with a letter")
	}

	if len(input.Password) < 8 || len(input.Password) > 64 {
		return input, fmt.Errorf("password must be between 8 and 64 characters")
	}

	hasLower := regexp.MustCompile(`[a-z]`).MatchString(input.Password)
	hasUpper := regexp.MustCompile(`[A-Z]`).MatchString(input.Password)
	hasNumber := regexp.MustCompile(`[0-9]`).MatchString(input.Password)
	hasSpecial := regexp.MustCompile(`[^a-zA-Z0-9]`).MatchString(input.Password)

	if !(hasLower && hasUpper && hasNumber && hasSpecial) {
		return input, fmt.Errorf("password must contain at least 1 uppercase, 1 lowercase, 1 number, and 1 special character")
	}

	return input, nil
}

func (h *userHandler) SignUp(ctx *gin.Context) {
	input, err := validateUserInput(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Use bcrypt to hash the password
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
		if strings.Contains(err.Error(), "Duplicate entry") {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Username already exists"})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "User created successfully"})
}

func (h *userHandler) Login(ctx *gin.Context) {
	input, err := validateUserInput(ctx)
	if err != nil {
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
