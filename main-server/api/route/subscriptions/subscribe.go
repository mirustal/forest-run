package subscriptions

import (
	"forest-run/main-server/api/controller/subscriptions"
	"forest-run/main-server/database"
	"forest-run/main-server/notifications"
	"github.com/gofiber/fiber/v2"
)

// InitSubscribe
// @Summary		Subscribe to user
// @Description	Start receiving notifications of user's actions
// @Tags			subscriptions
// @Produce		json
// @Security		ApiKeyAuth
// @Param			input	body		protocol.SubscriptionRequest	true "input"
// @Success		200		{object}	protocol.SubscriptionResponse
// @Failure		400,401	{object}	protocol.ErrorResponse
// @Failure		500		{object}	protocol.ErrorResponse
// @Router			/api/subscribe [post]
func InitSubscribe(group fiber.Router, notifs notifications.Manager, db database.DbAdapter) {
	c := subscriptions.NewSubscribe(notifs, db)
	group.Post("/subscribe", c.Handle)
}
