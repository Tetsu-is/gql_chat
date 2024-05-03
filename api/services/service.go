package services

import (
	"context"

	"github.com/Tetsu-is/gql_chat/graph/model"
	"gorm.io/gorm"
)

type Services interface {
	MessageService
	UserService
}

type services struct {
	*messageServices
	*userServices
}

type MessageService interface {
	GetLatestMessagesByIndex(ctx context.Context, index int) ([]*model.Message, error)
}

type UserService interface {
	GetUserByID(ctx context.Context, id string) (*model.User, error)
}

func New(db *gorm.DB) Services {
	return &services{
		messageServices: &messageServices{db: db},
		userServices:    &userServices{db: db},
	}
}
