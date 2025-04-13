package external

import (
	"context"
	"englishAI/config"
	"log"

	"google.golang.org/genai"
)

type aiService struct{}

type IAiService interface {
	Reply(message string) (string, error)
}

func NewAiService() IAiService {
	return &aiService{}
}

func (ai *aiService) Reply(message string) (string, error) {
	result, err := config.GetGenAiSer().SendMessage(context.Background(), genai.Part{Text: message})
	if err != nil {
		log.Fatal(err)
		return "AI has error, can not reply now! Please try later", err
	}
	return result.Text(), nil
}
