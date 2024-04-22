package runs

import (
	"github.com/gofiber/fiber/v2"
	"main-server/api/controller"
	"main-server/api/middleware"
	"main-server/database"
	"main-server/domain"
	"main-server/notifications"
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
	request := new(domain.InviteRunRequest)
	if err := ctx.BodyParser(request); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(domain.ErrorResponse{Message: "can't parse request json"})
	}

	authData := middleware.GetAuthData(ctx)

	run, err := c.db.GetRun(request.RunId, ctx.UserContext())
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(domain.ErrorResponse{Message: "can't get run from db"})
	}

	if run.Creator != authData.UserId {
		return ctx.Status(http.StatusMethodNotAllowed).JSON(domain.ErrorResponse{Message: "you can't invite this run"})
	}

	return ctx.Status(http.StatusNotImplemented).JSON(domain.InviteRunResponse{})
}
