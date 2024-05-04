package services

import (
	"context"
	"fmt"
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
		ID:    strconv.FormatUint(uint64(user.ID), 10),
		Name:  user.UserName,
		Email: user.Email,
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

func (s *userServices) CreateUser(ctx context.Context, input model.NewUser) (*model.User, error) {
	user := database.User{
		UserName: input.Name,
		Email:    input.Email,
	}
	result := s.db.Create(&user)

	if result.RowsAffected == 0 {
		return nil, fmt.Errorf("failed to insert user")
	}
	return convertUser(&user), nil
}

func (s *userServices) UpdateUser(ctx context.Context, id string, input model.UpdateUserInput) (*model.User, error) {
	if input.Name == nil && input.Email == nil {
		return nil, fmt.Errorf("no update input")
	}

	parsedId, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return nil, fmt.Errorf("failed to parse id")
	}
	var user database.User
	result := s.db.Where("id = ?", parsedId).First(&user)
	if result.RowsAffected == 0 {
		return nil, fmt.Errorf("user not found")
	}

	if input.Name != nil {
		user.UserName = *input.Name
	}
	if input.Email != nil {
		user.Email = *input.Email
	}

	result = s.db.Save(&user)
	if result.RowsAffected == 0 {
		return nil, fmt.Errorf("failed to update user")
	}
	return convertUser(&user), nil
}
