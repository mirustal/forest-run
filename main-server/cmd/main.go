package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/swagger"
	"go.uber.org/zap"
	"main-server/api/route"
	"main-server/boot"
	"main-server/database"
	_ "main-server/docs"

	"github.com/gofiber/fiber/v2/middleware/recover"
)

//	@title		Forest Run API
//	@version	0.0.1

//	@host		localhost:8080
//	@BasePath	/

//	@securityDefinitions.apiKey	ApiKeyAuth
//	@in							header
//	@name						Authorization

func main() {
	env, err := boot.NewEnv()
	if err != nil {
		panic(err)
	}

	logger := boot.NewLogger(env)
	defer logger.Sync()

	app := fiber.New()
	app.Use(recover.New())

	db, err := database.NewAdapter(env, logger)
	if err != nil {
		logger.Fatal("Error on initializing DB: ", zap.Error(err))
		panic(err)
	}

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept, Authorization",
	}))

	route.Setup(app, env, logger, db)

	if env.AppEnv == boot.DevEnv {
		app.Get("/swagger/*", swagger.HandlerDefault)
	}

	if err := app.Listen(env.ServerAddress); err != nil {
		logger.Fatal("Oops... Server is not running! Reason: ", zap.Error(err))
	}
}
