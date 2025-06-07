package repository

import (
	"englishAI/config"
	"englishAI/entities"
)

type IMessageRepo interface {
	Create(newMessage *entities.Message) error
	GetAll(conversationId string) ([]entities.Message, error)
}

type messageRepo struct{}

func NewMessageRepo() IMessageRepo {
	return &messageRepo{}
}

func (*messageRepo) Create(newMessage *entities.Message) error {
	return config.GetDatabase().Create(newMessage).Error
}

func (*messageRepo) GetAll(conversationId string) ([]entities.Message, error) {
	var messages []entities.Message
	err := config.GetDatabase().Where("conversation_id = ?", conversationId).Find(&messages).Error
	if err != nil {
		return nil, err
	}
	return messages, nil
}
