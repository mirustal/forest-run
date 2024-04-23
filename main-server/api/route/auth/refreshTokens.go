package auth

import (
	"forest-run/main-server/api/controller/auth"
	"forest-run/main-server/database"
	"forest-run/main-server/jwt"
	"github.com/gofiber/fiber/v2"
)

// InitRefreshTokens
// @Summary		Refresh Tokens
// @Description	Refresh JWT Tokens pair
// @Tags			auth
// @Produce		json
// @Security		ApiKeyAuth
// @Param			input	body		protocol.RefreshTokensRequest	true	"Refresh tokens data"
// @Success		200		{object}	protocol.RefreshTokensResponse
// @Failure		400,401	{object}	protocol.ErrorResponse
// @Failure		500		{object}	protocol.ErrorResponse
// @Router			/auth/refresh [post]
func InitRefreshTokens(group fiber.Router, jwt jwt.Provider, db database.DbAdapter) {
	c := auth.NewRefreshTokens(db, jwt)
	group.Post("/auth/refresh", c.Handle)
}
