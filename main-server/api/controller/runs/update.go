package runs

import (
	"forest-run/common/defs"
	"forest-run/common/middleware"
	"forest-run/main-server/api/controller"
	"forest-run/main-server/api/protocol"
	"forest-run/main-server/database"
	"forest-run/main-server/notifications"
	"forest-run/main-server/purchasing"
	"github.com/gofiber/fiber/v2"
	"net/http"
	"time"
)

type update struct {
	db        database.DbAdapter
	notifs    notifications.Manager
	defs      defs.Defs
	purchases purchasing.Manager
}

func NewUpdate(db database.DbAdapter, notifs notifications.Manager, defs defs.Defs, purchases purchasing.Manager) controller.Controller {
	return &update{db: db, notifs: notifs, defs: defs, purchases: purchases}
}

func (c update) Handle(ctx *fiber.Ctx) error {
	request := new(protocol.UpdateRunRequest)
	if err := ctx.BodyParser(request); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(protocol.ErrorResponse{Message: "can't parse request json"})
	}

	authData := middleware.GetAuthData(ctx)

	run, err := c.db.GetRun(request.RunId, ctx.UserContext())
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(protocol.ErrorResponse{Message: "can't get run from db"})
	}

	if run.Creator != authData.UserId {
		return ctx.Status(http.StatusMethodNotAllowed).JSON(protocol.ErrorResponse{Message: "you can't update this run"})
	}

	if request.PermissionsTransactionId != nil && request.RunPermissions != nil {
		err := c.purchases.ValidateRunPermissionsTransaction(*request.PermissionsTransactionId, *request.RunPermissions, ctx.UserContext())
		if err != nil {
			return ctx.Status(http.StatusBadRequest).JSON(protocol.ErrorResponse{Message: "invalid permissions transaction"})
		}

		run.RunPermissions = *request.RunPermissions
	}

	if request.StartTime != nil {
		if (*request.StartTime).Before(time.Now()) {
			return ctx.Status(http.StatusBadRequest).JSON(protocol.ErrorResponse{Message: "start time can't be in the past"})
		}

		run.StartTime = *request.StartTime
	}

	if request.RegistrationUntil != nil {
		if (*request.RegistrationUntil).Before(time.Now()) {
			return ctx.Status(http.StatusBadRequest).JSON(protocol.ErrorResponse{Message: "registration time can't be in the past"})
		}

		if (*request.RegistrationUntil).After(run.StartTime) {
			return ctx.Status(http.StatusBadRequest).JSON(protocol.ErrorResponse{Message: "registration time can't be after start time"})
		}

		run.RegistrationUntil = *request.RegistrationUntil
	}

	if request.Name != nil {
		run.Name = *request.Name
	}

	if request.Description != nil {
		run.Description = request.Description
	}

	if request.OfficialSiteUrl != nil {
		run.OfficialSiteUrl = request.OfficialSiteUrl
	}

	if request.AvatarUrl != nil {
		run.AvatarUrl = request.AvatarUrl
	}

	if request.StartPlace != nil {
		run.StartPlace = *request.StartPlace
	}

	if request.MaxParticipants != nil {
		run.MaxParticipants = *request.MaxParticipants
	}

	if request.ParticipationFormat != nil {
		run.ParticipationFormat = *request.ParticipationFormat
	}

	if request.Route != nil {
		run.Route = *request.Route
	}

	err = c.db.UpdateRun(run, ctx.UserContext())
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(protocol.ErrorResponse{Message: "can't store run"})
	}

	return ctx.Status(http.StatusOK).JSON(protocol.UpdateRunResponse{Run: run})
}
