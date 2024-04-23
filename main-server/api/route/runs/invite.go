package runs

import (
	"forest-run/main-server/api/controller/runs"
	"forest-run/main-server/database"
	"forest-run/main-server/notifications"
	"github.com/gofiber/fiber/v2"
)

// InitInvite
// @Summary		Invite run
// @Description	Invite run
// @Tags			runs
// @Produce		json
// @Security		ApiKeyAuth
// @Param			input	body		protocol.InviteRunRequest	true "input"
// @Success		200		{object}	protocol.InviteRunResponse
// @Failure		400,401,405	{object}	protocol.ErrorResponse
// @Failure		500		{object}	protocol.ErrorResponse
// @Router			/api/runs/Invite [post]
func InitInvite(group fiber.Router, db database.DbAdapter, notifs notifications.Manager) {
	c := runs.NewInvite(db, notifs)
	group.Post("/runs/invite", c.Handle)
}
