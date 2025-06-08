package usecase

import (
	"context"
	"englishAI/entities"
)

type IVocabularyCreateUsecase interface {
	Execute(ctx context.Context, name string, definition string, example string, pronunciation string, topicIds []uint, userID uint) (entities.Vocabulary, error)
}

type IVocabularyFindUsecase interface {
	Execute(ctx context.Context, paging entities.PagingRequest, userID uint) ([]entities.Vocabulary, uint32, error)
}

type IVocabularyDeleteUsecase interface {
	Execute(ctx context.Context, id string) error
}

type IVocabularyUpdateUsecase interface {
	Execute(ctx context.Context, id string, name string, definition string, example string, pronunciation string, topicIds []uint) (entities.Vocabulary, error)
}

// Conversation
type IConversationCreateUsecase interface {
	Execute(ctx context.Context, newConversation entities.Conversation) (entities.Conversation, error)
}

type IConversationFindUsecase interface {
	Execute(id string) (entities.Conversation, error)
}

// Topic
type ITopicCreateUsecase interface {
	Execute(ctx context.Context, topic entities.Topic) (entities.Topic, error)
}

type ITopicFindUsecase interface {
	Execute(ctx context.Context, paging entities.PagingRequest, userId uint) ([]entities.Topic, uint32, error)
}

type ITopicUpdateUsecase interface {
	Execute(ctx context.Context, id string, name string) (entities.Topic, error)
}

type ITopicDeleteUsecase interface {
	Execute(ctx context.Context, id string) error
}

type IUserCreateUsecase interface {
	Execute(ctx context.Context, user entities.User) error
}

type IUserFindUsecase interface {
	Execute(ctx context.Context, username string) (*entities.User, error)
}
type ITopicFindAllUsecase interface {
	Execute(userId int) ([]entities.Conversation, error)
}
