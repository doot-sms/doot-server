package services

import (
	"context"

	"github.com/doot-sms/doot-server/pkg/db"
)

type ISenderService interface {
	CreateSender(c context.Context, data CreateSenderParams) (db.Sender, error)
}

type SenderService struct {
	db *db.Queries
}

func NewSenderService(queries *db.Queries) *SenderService {
	return &SenderService{
		db: queries,
	}
}

type CreateSenderParams struct {
	UserId   int32  `json:"user_id"`
	DeviceId string `json:"device_id"`
}

func (SenderService *SenderService) CreateSender(c context.Context, data CreateSenderParams) (db.Sender, error) {
	args := db.CreateSenderParams{
		UserID:   data.UserId,
		DeviceID: data.DeviceId,
	}

	Sender, err := SenderService.db.CreateSender(c, args)

	if err != nil {
		return Sender, err
	}

	return Sender, nil
}
