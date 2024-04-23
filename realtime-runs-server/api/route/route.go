package route

import (
	. "forest-run/common/defs"
	"forest-run/realtime-runs-server/api/controller"
	. "forest-run/realtime-runs-server/boot"
	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App, env Env, defs Defs) {
	app.Post("/api/runs", controller.NewConnect().Handle)
}
