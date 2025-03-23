package main

import (
	"englishAI/config"
	"englishAI/migrations"
	"englishAI/routes"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	config.ConnectDatabase()
	migrations.Migrate()

	log.Println("Connected to the database!")

	r := gin.Default()
	routes.UserRoutes(r)

	r.Run(":8080")
}
