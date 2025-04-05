package main

import (
	"englishAI/config"
	"englishAI/routes"
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	log.Println("Starting server...")
	config.ConnectDatabase()
	// migrations.Migrate()

	log.Println("Connected to the database!")

	// Allow all origins
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE"},
		AllowHeaders:     []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
	}))

	routes.Routes(r)

	r.Run(":8080")
}
