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
	db *db.Queries
}

func NewAuthService(queries *db.Queries) *AuthService {
	return &AuthService{
		db: queries,
	}
}

var ErrInvalidCredentials = fmt.Errorf("invalid credentials")

func (authService *AuthService) Login(c context.Context, args LoginParams) (*LoginResponse, error) {

	config, err := utils.LoadConfig()

	if err != nil {
		return nil, fmt.Errorf("cannot load config: %w", err)
	}

	user, err := authService.db.GetUserByEmail(c, args.Email)

	if err != nil {
		return nil, ErrInvalidCredentials
	}

	err = utils.CheckPassword(args.Password, user.Password)
	if err != nil {
		return nil, ErrInvalidCredentials
	}

	symmetricKey := config.DootEncryptionKey

	pasetoMaker, err := token.NewPasetoMaker(symmetricKey)
	if err != nil {
		return nil, err
	}

	accessToken, tokenErr := pasetoMaker.CreateToken(user.ID, time.Minute*15)

	if tokenErr != nil {
		return nil, err
	}

	refreshToken, tokenErr := pasetoMaker.CreateToken(user.ID, time.Hour*24*7)

	if tokenErr != nil {
		return nil, err
	}

	return &LoginResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}
