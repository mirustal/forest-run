package route

import (
	"github.com/gofiber/fiber/v2"
	"main/boot"
)

func Setup(app *fiber.App, env *boot.Env) {
	InitHelloWorldRouter(app, env)
}
