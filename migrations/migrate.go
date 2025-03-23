package migrations

import (
	"englishAI/models"

	"englishAI/config"
	"fmt"
)

// Migrate runs the database migrations
func Migrate() {

	config.DB.AutoMigrate(&models.User{}, &models.Topic{}, &models.UserTopic{}, &models.Vocabulary{}, &models.TopicVocabulary{}, &models.UserVocabulary{}, &models.History{}, &models.Message{}, &models.HistoryMessage{})
	fmt.Println("Database migrated!")

}
