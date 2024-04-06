package route

import (
	"github.com/gofiber/fiber/v2"
	"main-server/api/controller"
	"main-server/boot"
)

// InitHelloWorldRouter
// @Summary      hello world
// @Description  hello world
// @Tags         helloworld
// @Produce      json
// @Success      200  {object}  domain.HelloWorldResponse
// @Router       /helloWorld [get]
func InitHelloWorldRouter(app *fiber.App, env *boot.Env) {
	c := controller.NewHelloWorldController()
	app.Get("/helloWorld", c.Handle)
}
