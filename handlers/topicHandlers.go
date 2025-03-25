package handlers

import (
	models "englishAI/entities"
	"englishAI/repository"

	"net/http"

	"github.com/gin-gonic/gin"
)

func GetTopics(c *gin.Context) {
	var topics []models.User

	// get topics from database
	if err := repository.GetDatabase().Find(&topics).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Cannot fetch topics"})
		return
	}

	// return  User list with JSON type
	c.JSON(http.StatusOK, topics)
}
