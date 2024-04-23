package runs

import (
	"forest-run/common/defs"
	"forest-run/main-server/api/controller/runs"
	"forest-run/main-server/database"
	"forest-run/main-server/notifications"
	"forest-run/main-server/purchasing"
	"github.com/gofiber/fiber/v2"
)

// InitUpdate
// @Summary		Update run
// @Description	Update run
// @Tags			runs
// @Produce		json
// @Security		ApiKeyAuth
// @Param			input	body		protocol.UpdateRunRequest	true "input"
// @Success		200		{object}	protocol.UpdateRunResponse
// @Failure		400,401,405	{object}	protocol.ErrorResponse
// @Failure		500		{object}	protocol.ErrorResponse
// @Router			/api/runs/update [post]
func InitUpdate(group fiber.Router, db database.DbAdapter, notifs notifications.Manager, defs defs.Defs, purchases purchasing.Manager) {
	c := runs.NewUpdate(db, notifs, defs, purchases)
	group.Post("/runs/update", c.Handle)
}
