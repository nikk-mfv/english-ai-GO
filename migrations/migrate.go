package migrations

import (
	"englishAI/entities"

	"englishAI/config"
	"fmt"
)

// Migrate runs the database migrations
func Migrate() {

	config.GetDatabase().AutoMigrate(&entities.User{}, &entities.Topic{}, &entities.Vocabulary{}, &entities.TopicVocabulary{}, &entities.History{}, &entities.Message{})
	fmt.Println("Database migrated!")
}
