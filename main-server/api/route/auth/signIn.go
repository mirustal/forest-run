package auth

import (
	"forest-run/common/jwt"
	"forest-run/main-server/api/controller/auth"
	"forest-run/main-server/database"
	"github.com/gofiber/fiber/v2"
)

// InitSignIn
// @Summary		Sign In
// @Description	Log into app and get access tokens
// @Tags			auth
// @Accepts		json
// @Produce		json
// @Param			input	body		protocol.SignInRequest	true	"SignIn data"
// @Success		200		{object}	protocol.SignInResponse
// @Failure		400,401	{object}	protocol.ErrorResponse
// @Failure		500		{object}	protocol.ErrorResponse
// @Router			/auth/sign-in [post]
func InitSignIn(app fiber.Router, db database.DbAdapter, jwt jwt.Provider) {
	c := auth.NewSignIn(db, jwt)
	app.Post("/auth/sign-in", c.Handle)
}
