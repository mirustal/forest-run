package subscriptions

import (
	"github.com/gofiber/fiber/v2"
	"main-server/api/controller/subscriptions"
	"main-server/database"
	"main-server/notifications"
)

func InitSubscribe(group fiber.Router, notifs notifications.Manager, db database.DbAdapter) {
	c := subscriptions.NewSubscribe(notifs, db)
	group.Post("/subscribe", c.Handle)
}
