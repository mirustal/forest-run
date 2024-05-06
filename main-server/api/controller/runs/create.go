package runs

import (
	"forest-run/common/defs"
	"forest-run/common/middleware"
	"forest-run/common/runs"
	"forest-run/main-server/api/controller"
	"forest-run/main-server/api/protocol"
	"forest-run/main-server/database"
	"forest-run/main-server/domain"
	"forest-run/main-server/notifications"
	"forest-run/main-server/purchasing"
	"github.com/gofiber/fiber/v2"
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
	request := new(protocol.CreateRunRequest)
	if err := ctx.BodyParser(request); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(protocol.ErrorResponse{Message: "can't parse request json"})
	}

	authData := middleware.GetAuthData(ctx)

	if request.PermissionsTransactionId != nil {
		err := c.purchases.ValidateRunPermissionsTransaction(*request.PermissionsTransactionId, request.RunPermissions, ctx.UserContext())
		if err != nil {
			return ctx.Status(http.StatusBadRequest).JSON(protocol.ErrorResponse{Message: "invalid permissions transaction"})
		}
	} else {
		request.RunPermissions = runs.FreePermissionsType
	}

	_, ok := c.defs.RunPermissionsDefs.Types[request.RunPermissions]
	if !ok {
		return ctx.Status(http.StatusBadRequest).JSON(protocol.ErrorResponse{Message: "invalid run permissions"})
	}

	if request.StartTime.Before(time.Now()) {
		return ctx.Status(http.StatusBadRequest).JSON(protocol.ErrorResponse{Message: "start time can't be in the past"})
	}

	run := domain.Run{
		Name:                request.Name,
		Creator:             authData.UserId,
		Description:         request.Description,
		OfficialSiteUrl:     request.OfficialSiteUrl,
		AvatarUrl:           request.AvatarUrl,
		MaxParticipants:     request.MaxParticipants,
		Status:              domain.OpenRunStatus,
		ParticipationFormat: request.ParticipationFormat,

		EssentialInfo: runs.EssentialInfo{
			Route:          request.Route,
			StartTime:      request.StartTime,
			StartPlace:     request.StartPlace,
			RunPermissions: request.RunPermissions,
		},
	}

	if request.RegistrationUntil != nil {
		run.RegistrationUntil = *request.RegistrationUntil
	} else {
		run.RegistrationUntil = run.StartTime
	}

	run, err := c.db.StoreRun(run, ctx.UserContext())
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(protocol.ErrorResponse{Message: "can't store run"})
	}

	notif, err := domain.Notification{
		Type:      notifications.NewRunCreated,
		CreatedAt: time.Now(),
	}.WithBody(notifications.RunCreatedBody{RunId: run.Id})

	if err != nil {
		ctx.Context().Logger().Printf("can't create notification body: ", err)
	} else {
		err = c.notifs.SendToSubscribers(authData.UserId, notif, ctx.UserContext())
		if err != nil {
			ctx.Context().Logger().Printf("error while sending run created notification: ", err)
		}
	}

	return ctx.Status(http.StatusOK).JSON(protocol.CreateRunResponse{Run: run})
}
