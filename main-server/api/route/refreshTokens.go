package route

import (
	"github.com/gofiber/fiber/v2"
	"main-server/api/controller"
	"main-server/database"
	"main-server/jwt"
)

// @Summary		Refresh Tokens
// @Description	Refresh JWT Tokens pair
// @Tags			auth
// @Produce		json
// @Security		ApiKeyAuth
// @Param			input	body		domain.RefreshTokensRequest	true	"Refresh tokens data"
// @Success		200		{object}	domain.RefreshTokensResponse
// @Failure		400,401	{object}	domain.ErrorResponse
// @Failure		500		{object}	domain.ErrorResponse
// @Router			/api/refresh [post]
func initRefreshTokens(group fiber.Router, jwt jwt.Provider, db database.DbAdapter) {
	c := controller.NewRefreshTokens(db, jwt)
	group.Post("/refresh", c.Handle)
}
