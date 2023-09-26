package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/doot-sms/doot-server/app/controllers"
	"github.com/doot-sms/doot-server/app/services"
	"github.com/doot-sms/doot-server/pkg/db"
	"github.com/doot-sms/doot-server/pkg/middlewares"
	"github.com/doot-sms/doot-server/pkg/routes"
	"github.com/doot-sms/doot-server/pkg/token"
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
	// config := configs.FiberConfig()

	config, err := utils.LoadConfig()

	if err != nil {
		log.Fatal("cannot load config: ", err)
		return
	}

	dbUrl := fmt.Sprintf("postgresql://%s:%s@db:5432/%s", config.DBUser, config.DBPassword, config.DBName)

	fmt.Println(dbUrl)

	database, err := sql.Open("postgres",
		dbUrl+"?sslmode=disable",
	)

	repository := db.New(database)

	if err != nil {
		log.Fatal(err)
		return
	}

	// paseto
	symmetricKey := config.DootEncryptionKey

	pasetoMaker, err := token.NewPasetoMaker(symmetricKey)

	// Define a new Fiber app with config.
	app := fiber.New(fiber.Config{
		ReadTimeout: time.Second * time.Duration(config.ServerReadTimeout),
	})

	// Middlewares.
	middlewares.FiberMiddleware(app) // Register Fiber's middleware for app.

	// services
	userService := services.NewUserService(repository)
	senderService := services.NewSenderService(repository)
	authService := services.NewAuthService(repository, config, pasetoMaker)

	// routes.SwaggerRoute(app)          // Register a route for API Docs (Swagger).

	controllers.ConnectUserRoutes(app, userService)
	controllers.ConnectSenderRoutes(app, senderService)

	controllers.ConnectAuthRoutes(app, authService)

	routes.NotFoundRoute(app) // Register route for 404 Error.

	// Start server (with or without graceful shutdown).
	if config.StageStatus == "dev" {
		utils.StartServer(app)
	} else {
		utils.StartServerWithGracefulShutdown(app)
	}
}
