package runs

import (
	"forest-run/common/defs"
	"forest-run/main-server/api/controller/runs"
	"forest-run/main-server/database"
	"forest-run/main-server/notifications"
	"forest-run/main-server/purchasing"
	"github.com/gofiber/fiber/v2"
)

// InitCreate
// @Summary		Create run
// @Description	Create run
// @Tags			runs
// @Produce		json
// @Security		ApiKeyAuth
// @Param			input	body		protocol.CreateRunRequest	true "input"
// @Success		200		{object}	protocol.CreateRunResponse
// @Failure		400,401	{object}	protocol.ErrorResponse
// @Failure		500		{object}	protocol.ErrorResponse
// @Router			/api/runs/create [post]
func InitCreate(group fiber.Router, db database.DbAdapter, notifs notifications.Manager, defs defs.Defs, purchases purchasing.Manager) {
	c := runs.NewCreate(db, notifs, defs, purchases)
	group.Post("/runs/create", c.Handle)
}
