package runs

import (
	"github.com/gofiber/fiber/v2"
	"main-server/api/controller/runs"
	"main-server/database"
	"main-server/notifications"
)

// InitInvite
// @Summary		Invite run
// @Description	Invite run
// @Tags			runs
// @Produce		json
// @Security		ApiKeyAuth
// @Param			input	body		domain.InviteRunRequest	true "input"
// @Success		200		{object}	domain.InviteRunResponse
// @Failure		400,401,405	{object}	domain.ErrorResponse
// @Failure		500		{object}	domain.ErrorResponse
// @Router			/api/runs/Invite [post]
func InitInvite(group fiber.Router, db database.DbAdapter, notifs notifications.Manager) {
	c := runs.NewInvite(db, notifs)
	group.Post("/runs/invite", c.Handle)
}
