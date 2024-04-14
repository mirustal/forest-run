package main

import (
	"github.com/gofiber/contrib/fiberzap/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/swagger"
	"go.uber.org/zap"
	"main-server/api/route"
	"main-server/boot"
	"main-server/database"
	_ "main-server/docs"
	"main-server/jwt"
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

	jwtProvider, err := jwt.NewProvider(env.JWTConfig)
	if err != nil {
		panic(err)
	}

	app := fiber.New()
	app.Use(recover.New())
	app.Use(fiberzap.New(fiberzap.Config{Logger: logger}))

	db, err := database.NewAdapter(env.DBConfig, logger)
	if err != nil {
		logger.Fatal("Error on initializing DB: ", zap.Error(err))
		panic(err)
	}

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept, Authorization",
	}))

	route.Setup(app, env, logger, db, jwtProvider)

	if env.AppEnv == boot.DevEnv {
		app.Get("/swagger/*", swagger.HandlerDefault)
	}

	if err := app.Listen(env.ServerAddress); err != nil {
		logger.Fatal("Oops... Server is not running! Reason: ", zap.Error(err))
	}
}
