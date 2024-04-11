package route

import (
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
	"main-server/api/controller"
	"main-server/database"
	"main-server/utils"
)

// signUp
//
//	@Summary		Sign In
//	@Description	Log into app and get access tokens
//	@Tags			auth
//	@Accepts		json
//	@Produce		json
//	@Param			input	body		domain.SignInRequest	true	"SignIn data"
//	@Success		200		{object}	domain.SignInResponse
//	@Failure		400,401	{object}	domain.ErrorResponse
//	@Failure		500		{object}	domain.ErrorResponse
//	@Router			/auth/sign-in [post]
func initSignIn(app *fiber.App, db database.DbAdapter, logger *zap.Logger, jwt utils.JWTProvider) {
	c := controller.NewSignIn(db, logger)
	app.Post("/auth/sign-in", c.Handle)
}
