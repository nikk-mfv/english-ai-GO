package main

import (
	"englishAI/migrations"
	"englishAI/repository"
	"englishAI/routes"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	repository.ConnectDatabase()
	migrations.Migrate()

	log.Println("Connected to the database!")

	r := gin.Default()
	routes.Routes(r)

	r.Run(":8080")
}
