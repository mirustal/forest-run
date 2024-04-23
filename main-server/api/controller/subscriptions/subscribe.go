package subscriptions

import (
	"forest-run/main-server/api/controller"
	"forest-run/main-server/api/middleware"
	"forest-run/main-server/api/protocol"
	"forest-run/main-server/database"
	"forest-run/main-server/domain"
	"forest-run/main-server/notifications"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

type subscribe struct {
	notifs notifications.Manager
	db     database.DbAdapter
}

func NewSubscribe(notifs notifications.Manager, db database.DbAdapter) controller.Controller {
	return &subscribe{notifs: notifs, db: db}
}

func (s subscribe) Handle(ctx *fiber.Ctx) error {
	request := new(protocol.SubscriptionRequest)
	if err := ctx.BodyParser(request); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(protocol.ErrorResponse{Message: "can't parse request json"})
	}

	authData := middleware.GetAuthData(ctx)
	subbed, err := s.db.Subscribe(authData.UserId, request.UserId, ctx.UserContext())
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(protocol.ErrorResponse{Message: "error while saving subscription to db"})
	}

	if subbed {
		notif, err := domain.Notification{
			FromUser: authData.UserId,
			ToUser:   request.UserId,
			Type:     notifications.NewSubscriber,
		}.WithBody(notifications.EmptyBody{})

		if err != nil {
			ctx.Context().Logger().Printf("error while sending subscription notification: ", err)
		} else {
			err = s.notifs.Send(notif, ctx.UserContext())
			if err != nil {
				ctx.Context().Logger().Printf("error while sending subscription notification: ", err)
			}
		}
	}

	return ctx.Status(http.StatusOK).JSON(protocol.SubscriptionResponse{})
}
