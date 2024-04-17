package runs

import (
	"github.com/gofiber/fiber/v2"
	"main-server/api/controller/runs"
	"main-server/database"
	"main-server/defs"
	"main-server/notifications"
	"main-server/purchasing"
)

// InitCreate
// @Summary		Create run
// @Description	Create run
// @Tags			runs
// @Produce		json
// @Security		ApiKeyAuth
// @Param			input	body		domain.CreateRunRequest	true "input"
// @Success		200		{object}	domain.CreateRunResponse
// @Failure		400,401	{object}	domain.ErrorResponse
// @Failure		500		{object}	domain.ErrorResponse
// @Router			/api/runs/create [post]
func InitCreate(group fiber.Router, db database.DbAdapter, notifs notifications.Manager, defs defs.Defs, purchases purchasing.Manager) {
	c := runs.NewCreate(db, notifs, defs, purchases)
	group.Post("/runs/create", c.Handle)
}
