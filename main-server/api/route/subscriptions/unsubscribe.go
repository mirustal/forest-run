package subscriptions

import (
	"forest-run/main-server/api/controller/subscriptions"
	"forest-run/main-server/database"
	"github.com/gofiber/fiber/v2"
)

// InitUnsubscribe
// @Summary		Unsubscribe from user
// @Description	Stop receiving notifications of user's actions
// @Tags			subscriptions
// @Produce		json
// @Security		ApiKeyAuth
// @Param			input	body		protocol.SubscriptionRequest	true "input"
// @Success		200		{object}	protocol.SubscriptionResponse
// @Failure		400,401	{object}	protocol.ErrorResponse
// @Failure		500		{object}	protocol.ErrorResponse
// @Router			/api/unsubscribe [post]
func InitUnsubscribe(group fiber.Router, db database.DbAdapter) {
	c := subscriptions.NewUnsubscribe(db)
	group.Post("/subscribe", c.Handle)
}
