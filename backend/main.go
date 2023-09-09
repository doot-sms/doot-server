package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/doot-sms/doot-server/app/controllers"
	"github.com/doot-sms/doot-server/app/services"
	"github.com/doot-sms/doot-server/pkg/configs"
	"github.com/doot-sms/doot-server/pkg/db"
	"github.com/doot-sms/doot-server/pkg/middleware"
	"github.com/doot-sms/doot-server/pkg/routes"
	"github.com/doot-sms/doot-server/pkg/utils"

	_ "github.com/lib/pq"

	"github.com/gofiber/fiber/v2"

	_ "github.com/doot-sms/doot-server/docs" // load API Docs files (Swagger)

	_ "github.com/joho/godotenv/autoload" // load .env file automatically
)

// @title API
// @version 1.0
// @description This is an auto-generated API Docs.
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email your@mail.com
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @BasePath /api
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	// Define Fiber config.
	config := configs.FiberConfig()

	database, err := sql.Open("postgres",
		os.Getenv("DATABASE_URL")+"?sslmode=disable",
	)

	repository := db.New(database)

	if err != nil {
		log.Fatal(err)
		return
	}

	// Define a new Fiber app with config.
	app := fiber.New(config)

	// Middlewares.
	middleware.FiberMiddleware(app) // Register Fiber's middleware for app.

	// services
	userService := services.NewUserService(repository)

	// routes.SwaggerRoute(app)          // Register a route for API Docs (Swagger).
	// routes.PublicRoutes(app, queries) // Register a public routes for app.
	// routes.PrivateRoutes(app)         // Register a private routes for app.
	// routes.NotFoundRoute(app)         // Register route for 404 Error.

	controllers.ConnectUserRoutes(app, userService)
	routes.SenderRoutes(app, repository)

	// Start server (with or without graceful shutdown).
	if os.Getenv("STAGE_STATUS") == "dev" {
		utils.StartServer(app)
	} else {
		utils.StartServerWithGracefulShutdown(app)
	}
}
