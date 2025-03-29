package repository

import (
	"englishAI/config"
	"englishAI/entities"

	"net/http"

	"github.com/gin-gonic/gin"
)

type TopicRepository struct{}

func (r *TopicRepository) GetTopics(c *gin.Context) {
	var topics []entities.Topic

	//get topics from database
	if err := config.GetDatabase().Find(&topics).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Cannot fetch topics"})
		return
	}

	//return topics list
	c.JSON(http.StatusOK, topics)
}

func (r *TopicRepository) CreateTopics(c *gin.Context) {
	var newTopic entities.Topic

	// get topics from database
	if err := c.ShouldBindJSON(&newTopic); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	// add new topic to database
	if err := config.GetDatabase().Create(&newTopic).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Cannot create topic"})
	}

	//return new topic
	c.JSON(http.StatusCreated, newTopic)
}
