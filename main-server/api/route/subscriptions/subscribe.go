package subscriptions

import (
	"github.com/gofiber/fiber/v2"
	"main-server/database"
	"main-server/notifications"
)

type subscribe struct {
	notifs notifications.Manager
	db     database.DbAdapter
}

func newSubscribe(notifs notifications.Manager, db database.DbAdapter) *subscribe {
	return &subscribe{notifs: notifs, db: db}
}

func (s subscribe) Handle(ctx *fiber.Ctx) error {

}
