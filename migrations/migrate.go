package migrations

import (
	"englishAI/entities"

	"englishAI/config"
	"fmt"
)

// func RenameHistoryTable() {
// 	db := config.GetDatabase()
// 	// Rename the table
// 	db.Exec("ALTER TABLE histories RENAME TO conversations")
// 	// If you have foreign keys, you might need to rename those constraints too
// 	// db.Exec("ALTER TABLE messages RENAME CONSTRAINT fk_histories_messages TO fk_conversations_messages")
// }

// Migrate runs the database migrations
func Migrate() {
	// RenameHistoryTable()
	config.GetDatabase().AutoMigrate(&entities.User{}, &entities.Topic{}, &entities.Vocabulary{}, &entities.Conversation{}, &entities.Message{})
	fmt.Println("Database migrated!")
}
