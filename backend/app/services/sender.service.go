package services

import (
	"context"

	"github.com/doot-sms/doot-server/pkg/db"
)

type ISenderService interface {
	CreateSender(c context.Context, data CreateSenderParams) (db.Sender, error)
	GetSenderByID(c context.Context, senderId int32) (db.Sender, error)
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
	UserId   int32
	DeviceId string
}

func (senderService *SenderService) CreateSender(c context.Context, data CreateSenderParams) (db.Sender, error) {
	args := db.CreateSenderParams{
		UserID:   data.UserId,
		DeviceID: data.DeviceId,
	}

	Sender, err := senderService.db.CreateSender(c, args)

	if err != nil {
		return Sender, err
	}

	return Sender, nil
}

func (senderService *SenderService) GetSenderByID(c context.Context, senderId int32) (db.Sender, error) {
	sender, err := senderService.db.GetSenderByID(c, senderId)

	if err != nil {
		return sender, err
	}

	return sender, nil
}
