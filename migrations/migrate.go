package migrations

import (
	"englishAI/entities"

	"englishAI/repository"
	"fmt"
)

// Migrate runs the database migrations
func Migrate() {

	repository.GetDatabase().AutoMigrate(&entities.User{}, &entities.Topic{}, &entities.Vocabulary{}, &entities.TopicVocabulary{}, &entities.History{}, &entities.Message{})
	fmt.Println("Database migrated!")
}
