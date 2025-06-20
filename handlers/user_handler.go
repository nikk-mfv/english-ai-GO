package handlers

import (
	"bytes"
	"context"
	"englishAI/config"
	"englishAI/entities"
	"englishAI/repository"
	"englishAI/usecase"
	"englishAI/utils"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type userHandler struct {
	BaseHandler
}

var userRepo = repository.NewUserRepository()
var ucUserCreate = usecase.NewUserCreateUsecase(userRepo)
var ucUserFind = usecase.NewUserFindUsecase(userRepo)
var ucUserUploadAvatar = usecase.NewUserUploadAvatarUsecase(userRepo)

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

// Content-Type file extension
func getContentTypeByExt(ext string) string {
	switch ext {
	case ".jpg", ".jpeg":
		return "image/jpeg"
	case ".png":
		return "image/png"
	default:
		return "application/octet-stream"
	}
}

func (h *userHandler) UploadAvatar(ctx *gin.Context) {
	userID, ok := ctx.Get("user_id")
	if !ok {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	// open file from form-data
	file, err := ctx.FormFile("avatar")
	if err != nil {

		ctx.JSON(http.StatusBadRequest, gin.H{"error": "No file uploaded"})
		return
	}

	// Check file size(5MB limit)
	if file.Size > 5*1024*1024 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "File size exceeds 5MB"})
		return
	}

	// Check format (allow jpg, jpeg, png)
	ext := filepath.Ext(file.Filename)
	allowedExt := map[string]bool{".jpg": true, ".jpeg": true, ".png": true}
	if !allowedExt[ext] {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid file type. Only jpg, jpeg, png allowed"})
		return
	}

	// Open file
	src, err := file.Open()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error opening file"})
		return
	}
	defer src.Close()

	// read all file in memory
	fileBytes, err := io.ReadAll(src)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error reading file"})
		return
	}

	// create S3 key
	fileKey := fmt.Sprintf("avatars/%d_%d%s", time.Now().Unix(), userID, ext)

	// Upload to S3
	s3cli := config.GetS3Client()
	bucket := os.Getenv("AVATAR_BUCKET")
	if bucket == "" {
		bucket = "avatar-bucket"
	}
	fmt.Println("Using bucket:", bucket)
	_, err = s3cli.PutObject(context.TODO(), &s3.PutObjectInput{
		Bucket:      &bucket,
		Key:         &fileKey,
		Body:        bytes.NewReader(fileBytes),
		ContentType: aws.String(getContentTypeByExt(ext)),
	})
	if err != nil {
		fmt.Println("S3 Upload error:", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Upload to S3 failed"})
		return
	}

	avatarURL := fmt.Sprintf("http://localhost:4566/%s/%s", bucket, fileKey)

	// Save avatar URL to user profile
	uid, ok := userID.(uint)
	if !ok {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid user ID type"})
		return
	}
	err = ucUserUploadAvatar.Execute(ctx, uid, avatarURL)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error saving avatar URL"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message":   "Avatar uploaded successfully",
		"image_url": avatarURL,
	})
}

func (h *userHandler) Profile(ctx *gin.Context) {
	userID := ctx.MustGet("user_id").(uint)
	username := ctx.MustGet("username").(string)
	imageURL, ok := ctx.Get("image_url")
	if !ok {
		imageURL = nil
	}

	ctx.JSON(http.StatusOK, gin.H{
		"user_id":   userID,
		"username":  username,
		"image_url": imageURL,
	})
}
