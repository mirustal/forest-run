package runs

import (
	"github.com/gofiber/fiber/v2"
	"main-server/api/controller/runs"
	"main-server/database"
	"main-server/defs"
	"main-server/notifications"
	"main-server/purchasing"
)

// InitUpdate
// @Summary		Update run
// @Description	Update run
// @Tags			runs
// @Produce		json
// @Security		ApiKeyAuth
// @Param			input	body		domain.UpdateRunRequest	true "input"
// @Success		200		{object}	domain.UpdateRunResponse
// @Failure		400,401,405	{object}	domain.ErrorResponse
// @Failure		500		{object}	domain.ErrorResponse
// @Router			/api/runs/update [post]
func InitUpdate(group fiber.Router, db database.DbAdapter, notifs notifications.Manager, defs defs.Defs, purchases purchasing.Manager) {
	c := runs.NewUpdate(db, notifs, defs, purchases)
	group.Post("/runs/update", c.Handle)
}
