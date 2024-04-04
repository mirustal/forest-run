package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"go.uber.org/zap"
	"main/api/route"
	"main/boot"
)

func main() {
	env, err := boot.NewEnv()
	if err != nil {
		panic(err)
	}

	logger := boot.NewLogger(env)
	defer logger.Sync()

	app := fiber.New()
	db, err := boot.NewDB(env, logger)
	if err != nil {
		logger.Fatal("Error on initializing DB: ", zap.Error(err))
		panic(err)
	}

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept, Authorization",
	}))

	route.Setup(app, env, logger, db)

	if err := app.Listen(env.ServerAddress); err != nil {
		logger.Fatal("Oops... Server is not running! Reason: ", zap.Error(err))
	}
}
