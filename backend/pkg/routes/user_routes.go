package routes

import (
	"github.com/doot-sms/doot-server/app/controllers"
	"github.com/doot-sms/doot-server/pkg/db"
	"github.com/gofiber/fiber/v2"
)

// PublicRoutes func for describe group of public routes.
func UserRoutes(a *fiber.App, db *db.Queries) {
	// Create routes group.
	route := a.Group("/api/v1/users")
	route.Post("", controllers.NewUserController(db).UserRegister) // register a new sender
}
