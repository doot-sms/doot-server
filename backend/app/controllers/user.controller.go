package controllers

import (
	"github.com/doot-sms/doot-server/pkg/db"
	"github.com/gofiber/fiber/v2"
)

type UserController struct {
	db *db.Queries
}

func NewUserController(queries *db.Queries) *UserController {
	return &UserController{
		db: queries,
	}
}

type CreateUserRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (
	sc *UserController,
) UserRegister(c *fiber.Ctx) error {
	var req CreateUserRequest

	if err := c.BodyParser(&req); err != nil {
		return err
	}

	args := db.CreateUserParams{
		Email:    req.Email,
		Password: req.Password,
	}

	user, err := sc.db.CreateUser(c.Context(), args)
	if err != nil {
		return err
	}

	c.JSON(fiber.Map{
		"message": "success",
		"user":    user,
	})

	return nil
}
