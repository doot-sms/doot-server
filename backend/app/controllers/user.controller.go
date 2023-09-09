package controllers

import (
	"github.com/doot-sms/doot-server/app/services"
	"github.com/gofiber/fiber/v2"
)

type UserController struct {
	userService services.IUserService
}

func NewUserController(userService services.IUserService) *UserController {
	return &UserController{
		userService: userService,
	}
}

// PublicRoutes func for describe group of public routes.
func ConnectUserRoutes(a *fiber.App, userService services.IUserService) {
	// Create routes group.
	controller := NewUserController(userService)

	route := a.Group("/api/v1/users")
	route.Post("", controller.UserRegister)
}

type CreateUserRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (
	uc *UserController,
) UserRegister(c *fiber.Ctx) error {
	var req CreateUserRequest

	if err := c.BodyParser(&req); err != nil {
		return err
	}

	user, err := uc.userService.CreateUser(c.Context(), services.CreateUserParams{
		Email:    req.Email,
		Password: req.Password,
	})

	if err != nil {
		return err
	}

	c.JSON(fiber.Map{
		"message": "success",
		"user":    user,
	})

	return nil
}
