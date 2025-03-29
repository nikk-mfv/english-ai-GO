package main

import (
	"englishAI/config"
	"englishAI/migrations"
	"englishAI/routes"
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	config.ConnectDatabase()
	migrations.Migrate()

	log.Println("Connected to the database!")

	// Allow all origins
	r := gin.Default()
	r.Use(cors.Default())

	routes.Routes(r)

	r.Run(":8080")
}
