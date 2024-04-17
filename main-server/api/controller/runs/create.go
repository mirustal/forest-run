package runs

import (
	"github.com/gofiber/fiber/v2"
	"main-server/api/controller"
	"main-server/api/middleware"
	"main-server/database"
	"main-server/defs"
	"main-server/domain"
	"main-server/notifications"
	"main-server/purchasing"
	"net/http"
	"time"
)

type create struct {
	db        database.DbAdapter
	notifs    notifications.Manager
	defs      defs.Defs
	purchases purchasing.Manager
}

func NewCreate(db database.DbAdapter, notifs notifications.Manager, defs defs.Defs, purchases purchasing.Manager) controller.Controller {
	return &create{db: db, notifs: notifs, defs: defs, purchases: purchases}
}

func (c create) Handle(ctx *fiber.Ctx) error {
	request := new(domain.CreateRunRequest)
	if err := ctx.BodyParser(request); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(domain.ErrorResponse{Message: "can't parse request json"})
	}

	authData := middleware.GetAuthData(ctx)

	if request.PermissionsTransactionId != nil {
		perm, err := c.purchases.ValidateRunPermissionsTransaction(*request.PermissionsTransactionId)
		request.RunPermissions = perm
		if err != nil {
			return ctx.Status(http.StatusBadRequest).JSON(domain.ErrorResponse{Message: "invalid permissions transaction"})
		}
	} else {
		request.RunPermissions = domain.FreeRunPermissionsType
	}

	def, ok := c.defs.RunPermissionsDefs.Types[request.RunPermissions]
	if !ok {
		return ctx.Status(http.StatusBadRequest).JSON(domain.ErrorResponse{Message: "invalid run permissions"})
	}

	if request.StartTime.Before(time.Now()) {
		return ctx.Status(http.StatusBadRequest).JSON(domain.ErrorResponse{Message: "start time can't be in the past"})
	}

	run := domain.Run{
		Name:                request.Name,
		Creator:             authData.UserId,
		Description:         request.Description,
		OfficialSiteUrl:     request.OfficialSiteUrl,
		AvatarUrl:           request.AvatarUrl,
		Route:               request.Route,
		StartTime:           request.StartTime,
		StartPlace:          request.StartPlace,
		MaxParticipants:     min(def.MaxOnlineParticipants, request.MaxParticipants),
		RunPermissions:      request.RunPermissions,
		Status:              domain.OpenRunStatus,
		ParticipationFormat: request.ParticipationFormat,
	}

	if request.RegistrationUntil != nil {
		run.RegistrationUntil = *request.RegistrationUntil
	} else {
		run.RegistrationUntil = run.StartTime
	}

	run, err := c.db.StoreRun(run, ctx.UserContext())
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(domain.ErrorResponse{Message: "can't store run"})
	}

	return ctx.Status(http.StatusOK).JSON(domain.CreateRunResponse{Run: run})
}
