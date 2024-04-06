package route

import (
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
	"main-server/api/controller"
	"main-server/database"
)

// SignUp
//
//	@Summary		Sign Up
//	@Description	Register in app and get access token
//	@Tags			auth
//	@Accepts		json
//	@Produce		json
//	@Params			input body domain.SignUpRequest true "SignUp data"
//	@Success		200		{object}	domain.SignUpResponse
//	@Failure		400,409	{object}	domain.ErrorResponse
//	@Failure		500		{object}	domain.ErrorResponse
//	@Router			/auth/sign-up [post]
func initSignUp(app *fiber.App, db database.DbAdapter, logger *zap.Logger) {
	c := controller.NewSignUp(db, logger)
	app.Post("/auth/sign-up", c.Handle)
}
