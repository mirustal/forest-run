package subscriptions

import (
	"forest-run/main-server/api/controller"
	"forest-run/main-server/api/middleware"
	"forest-run/main-server/api/protocol"
	"forest-run/main-server/database"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

type unsubscribe struct {
	db database.DbAdapter
}

func NewUnsubscribe(db database.DbAdapter) controller.Controller {
	return &unsubscribe{db: db}
}

func (s unsubscribe) Handle(ctx *fiber.Ctx) error {
	request := new(protocol.SubscriptionRequest)
	if err := ctx.BodyParser(request); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(protocol.ErrorResponse{Message: "can't parse request json"})
	}

	authData := middleware.GetAuthData(ctx)

	err := s.db.Unsubscribe(authData.UserId, request.UserId, ctx.UserContext())
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(protocol.ErrorResponse{Message: "error while saving subscription to db"})
	}

	return ctx.Status(http.StatusOK).JSON(protocol.SubscriptionResponse{})
}
