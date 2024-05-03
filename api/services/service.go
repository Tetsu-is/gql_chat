package services

import (
	"context"

	"github.com/Tetsu-is/gql_chat/graph/model"
	"gorm.io/gorm"
)

type Services interface {
	MessageService
}

type services struct {
	*messageServices
}

type MessageService interface {
	GetLatestMessagesByIndex(ctx context.Context, index int) ([]*model.Message, error)
}

func New(db *gorm.DB) Services {
	return &services{
		messageServices: &messageServices{
			db: db,
		},
	}
}
