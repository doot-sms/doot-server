package services

import (
	"context"

	"github.com/doot-sms/doot-server/pkg/db"
	"github.com/doot-sms/doot-server/pkg/utils"
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

func (userService *UserService) CreateUser(c context.Context, args CreateUserParams) (db.User, error) {
	user, err := userService.db.CreateUser(c, db.CreateUserParams{
		Email:    args.Email,
		Password: utils.GeneratePassword(args.Password),
	})

	if err != nil {
		return user, err
	}

	return user, nil
}
