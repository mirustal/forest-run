package subscriptions

import (
	"github.com/gofiber/fiber/v2"
	"main-server/api/controller/subscriptions"
	"main-server/database"
)

// InitUnsubscribe
// @Summary		Unsubscribe from user
// @Description	Stop receiving notifications of user's actions
// @Tags			subscriptions
// @Produce		json
// @Security		ApiKeyAuth
// @Param			input	body		domain.SubscriptionRequest	true "input"
// @Success		200		{object}	domain.SubscriptionResponse
// @Failure		400,401	{object}	domain.ErrorResponse
// @Failure		500		{object}	domain.ErrorResponse
// @Router			/api/unsubscribe [post]
func InitUnsubscribe(group fiber.Router, db database.DbAdapter) {
	c := subscriptions.NewUnsubscribe(db)
	group.Post("/subscribe", c.Handle)
}
