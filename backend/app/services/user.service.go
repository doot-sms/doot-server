package services

import (
	"context"
	"fmt"
	"time"

	"github.com/doot-sms/doot-server/pkg/db"
	"github.com/doot-sms/doot-server/pkg/token"
	"github.com/doot-sms/doot-server/pkg/utils"
)

type LoginParams struct {
	Email    string
	Password string
}

type LoginResponse struct {
	AccessToken string
}

type IUserService interface {
	CreateUser(c context.Context, data CreateUserParams) (db.User, error)
	Login(c context.Context, data LoginParams) (*LoginResponse, error)
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

var ErrInvalidCredentials = fmt.Errorf("invalid credentials")

func (userService *UserService) Login(c context.Context, args LoginParams) (*LoginResponse, error) {

	user, err := userService.db.GetUserByEmail(c, args.Email)

	if err != nil {
		return nil, ErrInvalidCredentials
	}

	err = utils.CheckPassword(args.Password, user.Password)
	if err != nil {
		return nil, ErrInvalidCredentials
	}

	pasetoMaker, err := token.NewPasetoMaker("mR3N0Jm++FXvW/LuE7FfT3Y1C0nQlPNS")
	if err != nil {
		return nil, err
	}

	accessToken, err := pasetoMaker.CreateToken(user.ID, time.Minute*15)

	if err != nil {
		return nil, err
	}

	return &LoginResponse{
		AccessToken: accessToken,
	}, nil
}
