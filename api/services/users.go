package services

import (
	"context"
	"strconv"

	"github.com/Tetsu-is/gql_chat/database"
	"github.com/Tetsu-is/gql_chat/graph/model"
	"gorm.io/gorm"
)

type userServices struct {
	db *gorm.DB
}

func convertUser(user *database.User) *model.User {
	return &model.User{
		ID:   strconv.FormatUint(uint64(user.ID), 10),
		Name: user.UserName,
	}
}

func (s *userServices) GetUserByID(ctx context.Context, id string) (*model.User, error) {
	var user database.User
	s.db.Where("id = ?", id).First(&user)
	if user.ID == 0 {
		return nil, nil
	}
	return convertUser(&user), nil
}
