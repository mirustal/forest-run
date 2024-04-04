package route

import (
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
	"main/boot"
)

func Setup(app *fiber.App, env *boot.Env, logger *zap.Logger) {
	InitHelloWorldRouter(app, env)
}
