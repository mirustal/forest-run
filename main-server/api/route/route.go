package route

import (
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
	"main/boot"
	"main/database"
)

func Setup(app *fiber.App, env *boot.Env, logger *zap.Logger, db database.DbAdapter) {
	defer logger.Sync()
	InitHelloWorldRouter(app, env)
}
