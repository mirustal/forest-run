package auth

import (
	"github.com/gofiber/fiber/v2"
	"main-server/api/controller/auth"
	"main-server/database"
	"main-server/jwt"
)

// InitSignIn
// @Summary		Sign In
// @Description	Log into app and get access tokens
// @Tags			auth
// @Accepts		json
// @Produce		json
// @Param			input	body		domain.SignInRequest	true	"SignIn data"
// @Success		200		{object}	domain.SignInResponse
// @Failure		400,401	{object}	domain.ErrorResponse
// @Failure		500		{object}	domain.ErrorResponse
// @Router			/auth/sign-in [post]
func InitSignIn(app *fiber.App, db database.DbAdapter, jwt jwt.Provider) {
	c := auth.NewSignIn(db, jwt)
	app.Post("/auth/sign-in", c.Handle)
}
