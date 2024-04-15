package subscriptions

import (
	"github.com/gofiber/fiber/v2"
	"main-server/api/controller"
	"main-server/api/middleware"
	"main-server/database"
	"main-server/domain"
	"net/http"
)

type unsubscribe struct {
	db database.DbAdapter
}

func NewUnsubscribe(db database.DbAdapter) controller.Controller {
	return &unsubscribe{db: db}
}

func (s unsubscribe) Handle(ctx *fiber.Ctx) error {
	request := new(domain.SubscriptionRequest)
	if err := ctx.BodyParser(request); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(domain.ErrorResponse{Message: "can't parse request json"})
	}

	authData := middleware.GetAuthData(ctx)

	err := s.db.Unsubscribe(authData.UserId, request.UserId, ctx.UserContext())
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(domain.ErrorResponse{Message: "error while saving subscription to db"})
	}

	return ctx.Status(http.StatusOK).JSON(domain.SubscriptionResponse{})
}
