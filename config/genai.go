package config

import (
	"context"
	"log"
	"os"

	"github.com/joho/godotenv"
	"google.golang.org/genai"
)

var genAiSer *genai.Chat

func ConnectGenAi() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	ctx := context.Background()
	apiKey := os.Getenv("GEMINI_API_KEY")

	client, err := genai.NewClient(ctx, &genai.ClientConfig{
		APIKey:  apiKey,
		Backend: genai.BackendGeminiAPI,
	})
	if err != nil {
		log.Fatal(err)
	}

	chat, err := client.Chats.Create(ctx, "gemini-2.0-flash-exp", nil, nil)
	if err != nil {
		log.Fatal(err)
	}

	genAiSer = chat
}

func GetGenAiSer() *genai.Chat {
	if genAiSer == nil {
		log.Fatal("GenAI not initialized, call ConnectDatabase() first")
	}
	return genAiSer
}
