package controllers

import (
	"github.com/doot-sms/doot-server/app/services"
	"github.com/gofiber/fiber/v2"
)

type AuthController struct {
	authService services.IAuthService
}

func NewAuthController(authService services.IAuthService) *AuthController {
	return &AuthController{
		authService: authService,
	}
}

// PublicRoutes func for describe group of public routes.
func ConnectAuthRoutes(a *fiber.App, authService services.IAuthService) {
	// Create routes group.
	authController := NewAuthController(authService)

	route := a.Group("/api/v1/auth")
	route.Post("/tokens", authController.Login)
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (
	ac *AuthController,
) Login(c *fiber.Ctx) error {
	var req LoginRequest

	if err := c.BodyParser(&req); err != nil {
		return err
	}

	userTokens, err := ac.authService.Login(c.Context(), services.LoginParams{
		Email:    req.Email,
		Password: req.Password,
	})

	if err != nil {
		return err
	}

	c.JSON(fiber.Map{
		"message": "success",
		"user":    userTokens,
	})

	return nil
}
