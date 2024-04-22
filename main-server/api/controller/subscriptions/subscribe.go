package subscriptions

import (
	"github.com/gofiber/fiber/v2"
	"main-server/api/controller"
	"main-server/api/middleware"
	"main-server/database"
	"main-server/domain"
	"main-server/notifications"
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
	request := new(domain.SubscriptionRequest)
	if err := ctx.BodyParser(request); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(domain.ErrorResponse{Message: "can't parse request json"})
	}

	authData := middleware.GetAuthData(ctx)
	subbed, err := s.db.Subscribe(authData.UserId, request.UserId, ctx.UserContext())
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(domain.ErrorResponse{Message: "error while saving subscription to db"})
	}

	if subbed {
		notif, err := domain.Notification{
			FromUser: authData.UserId,
			ToUser:   request.UserId,
			Type:     domain.NewSubscriberNotification,
		}.WithBody(domain.EmptyNotificationBody{})

		if err != nil {
			ctx.Context().Logger().Printf("error while sending subscription notification: ", err)
		} else {
			err = s.notifs.Send(notif, ctx.UserContext())
			if err != nil {
				ctx.Context().Logger().Printf("error while sending subscription notification: ", err)
			}
		}
	}

	return ctx.Status(http.StatusOK).JSON(domain.SubscriptionResponse{})
}
