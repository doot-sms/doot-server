package services

import (
	"context"

	"github.com/doot-sms/doot-server/pkg/db"
)

type IUserService interface {
	CreateUser(c context.Context, data CreateUserParams) (db.User, error)
}

type UserService struct {
	db *db.Queries
}

func NewUserService(queries *db.Queries) *UserService {
	return &UserService{
		db: queries,
	}
}

type CreateUserParams struct {
	Email    string
	Password string
}

func (userService *UserService) CreateUser(c context.Context, data CreateUserParams) (db.User, error) {
	args := db.CreateUserParams{
		Email:    data.Email,
		Password: data.Password,
	}

	user, err := userService.db.CreateUser(c, args)

	if err != nil {
		return user, err
	}

	return user, nil
}
