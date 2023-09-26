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
	AccessToken  string
	RefreshToken string
}

type IAuthService interface {
	Login(c context.Context, data LoginParams) (*LoginResponse, error)
}

type AuthService struct {
	db         *db.Queries
	config     utils.Config
	tokenMaker token.TokenMaker
}

func NewAuthService(queries *db.Queries, config utils.Config, tokenMaker token.TokenMaker) *AuthService {
	return &AuthService{
		db:         queries,
		config:     config,
		tokenMaker: tokenMaker,
	}
}

var ErrInvalidCredentials = fmt.Errorf("invalid credentials")

func (authService *AuthService) Login(c context.Context, args LoginParams) (*LoginResponse, error) {

	user, err := authService.db.GetUserByEmail(c, args.Email)

	if err != nil {
		return nil, ErrInvalidCredentials
	}

	err = utils.CheckPassword(args.Password, user.Password)
	if err != nil {
		return nil, ErrInvalidCredentials
	}

	accessToken, tokenErr := authService.tokenMaker.CreateToken(user.ID, time.Minute*15)

	if tokenErr != nil {
		return nil, err
	}

	refreshToken, tokenErr := authService.tokenMaker.CreateToken(user.ID, time.Hour*24*7)

	if tokenErr != nil {
		return nil, err
	}

	return &LoginResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}
