package runs

import (
	"forest-run/common/middleware"
	"forest-run/main-server/api/controller"
	"forest-run/main-server/api/protocol"
	"forest-run/main-server/database"
	"forest-run/main-server/notifications"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

type invite struct {
	db     database.DbAdapter
	notifs notifications.Manager
}

func NewInvite(db database.DbAdapter, notifs notifications.Manager) controller.Controller {
	return &invite{db: db, notifs: notifs}
}

func (c invite) Handle(ctx *fiber.Ctx) error {
	request := new(protocol.InviteRunRequest)
	if err := ctx.BodyParser(request); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(protocol.ErrorResponse{Message: "can't parse request json"})
	}

	authData := middleware.GetAuthData(ctx)

	run, err := c.db.GetRun(request.RunId, ctx.UserContext())
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(protocol.ErrorResponse{Message: "can't get run from db"})
	}

	if run.Creator != authData.UserId {
		return ctx.Status(http.StatusMethodNotAllowed).JSON(protocol.ErrorResponse{Message: "you can't invite this run"})
	}

	return ctx.Status(http.StatusNotImplemented).JSON(protocol.InviteRunResponse{})
}
