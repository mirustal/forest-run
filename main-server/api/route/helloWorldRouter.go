package route

import (
	"github.com/gofiber/fiber/v2"
	"main/api/controller"
	"main/boot"
)

func InitHelloWorldRouter(app *fiber.App, env *boot.Env) {
	c := controller.NewHelloWorldController()
	app.Get("/helloWorld", c.Handle)
}
