package controllers

import (
	"github.com/doot-sms/doot-server/pkg/db"
	"github.com/gofiber/fiber/v2"
)

type SenderController struct {
	db *db.Queries
}

func NewSenderController(queries *db.Queries) *SenderController {
	return &SenderController{
		db: queries,
	}
}

type CreateSenderRequest struct {
	UserId   int32  `json:"user_id"`
	DeviceId string `json:"device_id"`
}

func (
	sc *SenderController,
) SenderRegister(c *fiber.Ctx) error {
	var req CreateSenderRequest

	if err := c.BodyParser(&req); err != nil {
		return err
	}

	args := db.CreateSenderParams{
		UserID:   req.UserId,
		DeviceID: req.DeviceId,
	}

	sc.db.CreateSender(c.Context(), args)
	return nil
}
