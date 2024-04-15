package subscriptions

import (
	"github.com/gofiber/fiber/v2"
	"main-server/api/controller"
	"main-server/database"
	"main-server/notifications"
)

type subscribe struct {
	notifs notifications.Manager
	db     database.DbAdapter
}

func NewSubscribe(notifs notifications.Manager, db database.DbAdapter) controller.Controller {
	return &subscribe{notifs: notifs, db: db}
}

func (s subscribe) Handle(ctx *fiber.Ctx) error {
	return ctx.Status(200).SendString("todo")
}
