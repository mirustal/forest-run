package auth

import (
	"github.com/gofiber/fiber/v2"
	"main-server/api/controller/auth"
	"main-server/database"
)

// InitSignUp
// @Summary		Sign Up
// @Description	Register in app
// @Tags			auth
// @Accepts		json
// @Produce		json
// @Param			input	body		domain.SignUpRequest	true	"SignUp data"
// @Success		200		{object}	domain.SignUpResponse
// @Failure		400,409	{object}	domain.ErrorResponse
// @Failure		500		{object}	domain.ErrorResponse
// @Router			/auth/sign-up [post]
func InitSignUp(app fiber.Router, db database.DbAdapter) {
	c := auth.NewSignUp(db)
	app.Post("/auth/sign-up", c.Handle)
}
