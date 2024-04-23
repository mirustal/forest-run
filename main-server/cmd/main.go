package main

import (
	"forest-run/main-server/api/route"
	"forest-run/main-server/boot"
	"forest-run/main-server/database"
	_ "forest-run/main-server/docs"
	"forest-run/main-server/jwt"
	"forest-run/main-server/notifications"
	"forest-run/main-server/purchasing"
	"github.com/gofiber/contrib/fiberzap/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/swagger"
	"go.uber.org/zap"
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

	logger.Sugar().Info("loaded env: ", env)

	defs, err := boot.LoadDefs(env)
	if err != nil {
		logger.Fatal("Error on loading defs: ", zap.Error(err))
	}

	jwtProvider := jwt.NewProvider(env.JWTConfig)

	app := fiber.New()
	app.Use(recover.New())
	app.Use(fiberzap.New(fiberzap.Config{Logger: logger}))

	db, err := database.NewAdapter(env.DBConfig, logger)
	if err != nil {
		logger.Fatal("Error on initializing DB: ", zap.Error(err))
	}

	notifs := notifications.NewManager(db)
	purchasingManager := purchasing.NewManager()

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept, Authorization",
	}))

	route.Setup(app, db, jwtProvider, notifs, defs, purchasingManager)

	if env.AppEnv == boot.DevEnv {
		app.Get("/swagger/*", swagger.HandlerDefault)
	}

	if err := app.Listen(env.ServerAddress); err != nil {
		logger.Fatal("Oops... Server is not running! Reason: ", zap.Error(err))
	}
}
