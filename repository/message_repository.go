package repository

import (
	"englishAI/config"
	"englishAI/entities"
)

type IMessageRepo interface {
	Create(newMessage *entities.Message) error
}

type messageRepo struct{}

func NewMessageRepo() IMessageRepo {
	return &messageRepo{}
}

func (*messageRepo) Create(newMessage *entities.Message) error {
	return config.GetDatabase().Create(newMessage).Error
}
