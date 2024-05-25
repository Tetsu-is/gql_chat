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
	CreateMessage(ctx context.Context, input model.NewMessage) (*model.Message, error)
	PickMessageFromChannel() <-chan *model.Message
}

type UserService interface {
	GetUserByID(ctx context.Context, id string) (*model.User, error)
	CreateUser(ctx context.Context, input model.NewUser) (*model.User, error)
	UpdateUser(ctx context.Context, id string, input model.UpdateUserInput) (*model.User, error)
}

func New(db *gorm.DB) Services {
	return &services{
		messageServices: &messageServices{db: db, ch: make(chan *model.Message)},
		userServices:    &userServices{db: db},
	}
}
