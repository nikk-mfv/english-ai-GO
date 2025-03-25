package handlers

import (
	models "englishAI/entities"
	"englishAI/repository"

	"net/http"

	"github.com/gin-gonic/gin"
)

func GetUsers(c *gin.Context) {
	var users []models.User

	// get User from database
	if err := repository.GetDatabase().Find(&users).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Cannot fetch users"})
		return
	}

	// return  User list with JSON type
	c.JSON(http.StatusOK, users)
}

func CreateUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	repository.GetDatabase().Create(&user)
	c.JSON(http.StatusOK, user)
}
