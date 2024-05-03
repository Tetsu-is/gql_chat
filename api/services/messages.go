package services

import (
	"context"
	"fmt"
	"strconv"

	"github.com/Tetsu-is/gql_chat/database"
	"github.com/Tetsu-is/gql_chat/graph/model"
	"gorm.io/gorm"
)

type messageServices struct {
	db *gorm.DB
}

func convertMessage(message *database.Message) *model.Message {
	return &model.Message{
		ID:       strconv.FormatUint(uint64(message.ID), 10),
		UserID:   strconv.FormatInt(int64(message.UserID), 10),
		UserName: message.UserName,
		Body:     message.Body,
	}
}

func (s *messageServices) GetLatestMessagesByIndex(ctx context.Context, index int) ([]*model.Message, error) {
	var messages []database.Message
	s.db.Where("id >= ?", index).Find(&messages)
	if len(messages) == 0 {
		return nil, fmt.Errorf("no messages found")
	}
	var convertedMessages []*model.Message
	for _, message := range messages {
		convertedMessages = append(convertedMessages, convertMessage(&message))
	}
	return convertedMessages, nil
}

// func (s *messageServices) CreateMessage(ctx context.Context, input model.NewMessage) (*model.Message, error) {
// 	parsedID, err := strconv.ParseInt(input.UserID, 10, 64)
// 	if err != nil {
// 		return nil, fmt.Errorf("failed to parse user ID: %w", err)
// 	}
// }
