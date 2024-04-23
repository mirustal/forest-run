package auth

import (
	"forest-run/main-server/api/controller/auth"
	"forest-run/main-server/database"
	"github.com/gofiber/fiber/v2"
)

// InitSignUp
// @Summary		Sign Up
// @Description	Register in app
// @Tags			auth
// @Accepts		json
// @Produce		json
// @Param			input	body		protocol.SignUpRequest	true	"SignUp data"
// @Success		200		{object}	protocol.SignUpResponse
// @Failure		400,409	{object}	protocol.ErrorResponse
// @Failure		500		{object}	protocol.ErrorResponse
// @Router			/auth/sign-up [post]
func InitSignUp(app fiber.Router, db database.DbAdapter) {
	c := auth.NewSignUp(db)
	app.Post("/auth/sign-up", c.Handle)
}
