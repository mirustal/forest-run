package route

import (
	. "forest-run/common/defs"
	"forest-run/common/jwt"
	"forest-run/common/middleware"
	"forest-run/realtime-runs-server/api/controller"
	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App, defs Defs, jwt jwt.Provider) {
	middleware.InitJwtAuth(app, jwt)
	middleware.InitWebsocket(app)
	app.Get("/api/connect", websocket.New(controller.NewConnect(defs).Handle))
}
