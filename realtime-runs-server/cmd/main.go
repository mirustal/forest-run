package main

import (
	defs2 "forest-run/common/defs"
	"forest-run/common/jwt"
	logger2 "forest-run/common/logger"
	"forest-run/realtime-runs-server/api/route"
	"forest-run/realtime-runs-server/boot"
	"github.com/gofiber/contrib/fiberzap/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"go.uber.org/zap"
)

func main() {
	env, err := boot.NewEnv()
	if err != nil {
		panic(err)
	}

	logger := logger2.New(env.AppEnv, env.LoggerConfig)
	defer logger.Sync()

	logger.Sugar().Info("loaded env: ", env)

	defs, err := defs2.Load(env.CommonConfig)
	if err != nil {
		logger.Fatal("Error on loading defs: ", zap.Error(err))
	}

	app := fiber.New()
	app.Use(recover.New())
	app.Use(fiberzap.New(fiberzap.Config{Logger: logger}))
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept, Authorization",
	}))

	route.Setup(app, defs, jwt.NewProvider(env.JWTConfig))

	if err := app.Listen(env.ServerAddress); err != nil {
		logger.Fatal("Oops... Server is not running! Reason: ", zap.Error(err))
	}
}
