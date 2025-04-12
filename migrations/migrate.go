package migrations

import (
	"englishAI/entities"

	"englishAI/config"
	"fmt"
)

// Migrate runs the database migrations
func Migrate() {
	err := config.GetDatabase().AutoMigrate(&entities.User{}, &entities.Topic{}, &entities.Vocabulary{}, &entities.Conversation{}, &entities.Message{})
	if err != nil {
		fmt.Println("Error during migration", err)
		return
	}
	fmt.Println("Database migrated!")
}
