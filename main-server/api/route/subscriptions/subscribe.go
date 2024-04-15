package subscriptions

import (
	"github.com/gofiber/fiber/v2"
	"main-server/api/controller/subscriptions"
	"main-server/database"
	"main-server/notifications"
)

// InitSubscribe
// @Summary		Subscribe to user
// @Description	Start receiving notifications of user's actions
// @Tags			subscriptions
// @Produce		json
// @Security		ApiKeyAuth
// @Param			input	body		domain.SubscriptionRequest	true "input"
// @Success		200		{object}	domain.SubscriptionResponse
// @Failure		400,401	{object}	domain.ErrorResponse
// @Failure		500		{object}	domain.ErrorResponse
// @Router			/api/subscribe [post]
func InitSubscribe(group fiber.Router, notifs notifications.Manager, db database.DbAdapter) {
	c := subscriptions.NewSubscribe(notifs, db)
	group.Post("/subscribe", c.Handle)
}