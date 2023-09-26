package controllers

import (
	"github.com/doot-sms/doot-server/app/services"
	"github.com/gofiber/fiber/v2"
)

type SenderController struct {
	SenderService services.ISenderService
}

func NewSenderController(SenderService services.ISenderService) *SenderController {
	return &SenderController{
		SenderService: SenderService,
	}
}

// PublicRoutes func for describe group of public routes.
func ConnectSenderRoutes(a *fiber.App, SenderService services.ISenderService, cookieUserAuthMiddleware fiber.Handler) {
	// Create routes group.
	controller := NewSenderController(SenderService)

	route := a.Group("/api/v1/Senders")
	route.Post("", cookieUserAuthMiddleware, controller.SenderRegister)
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

	Sender, err := sc.SenderService.CreateSender(c.Context(), services.CreateSenderParams{
		UserId:   req.UserId,
		DeviceId: req.DeviceId,
	})

	if err != nil {
		return err
	}

	c.JSON(fiber.Map{
		"message": "success",
		"Sender":  Sender,
	})

	return nil
}

type GetSenderByIDRequestParams struct {
	Id int32 `json:"id"`
}

func (sc *SenderController) GetSenderByID(c *fiber.Ctx) error {
	var req GetSenderByIDRequestParams

	if err := c.BodyParser(&req); err != nil {
		return err
	}

	Sender, err := sc.SenderService.GetSenderByID(c.Context(), req.Id)

	if err != nil {
		return err
	}

	c.JSON(fiber.Map{
		"message": "success",
		"Sender":  Sender,
	})

	return nil
}
