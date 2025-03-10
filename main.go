package main

import (
	"english-ai-be/db"
	"english-ai-be/internal/handlers"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db.Connect()
	handlers.Router()
}
