package repository

import (
	"context"
	"englishAI/config"
	"englishAI/entities"
)

type conversationRepository struct{}

func NewConversationRepository() IConversationRepository {
	return &conversationRepository{}
}

func (r *conversationRepository) Create(ctx context.Context, obj *entities.Conversation) error {
	return config.GetDatabase().Create(obj).Error
}

func (r *conversationRepository) Find(id string) (entities.Conversation, error) {
	var conversation entities.Conversation
	err := config.GetDatabase().Preload("User").First(&conversation, id).Error
	return conversation, err
}
