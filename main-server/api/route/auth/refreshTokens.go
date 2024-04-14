package auth

import (
	"github.com/gofiber/fiber/v2"
	"main-server/api/controller/auth"
	"main-server/database"
	"main-server/jwt"
)

// InitRefreshTokens
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
func InitRefreshTokens(group fiber.Router, jwt jwt.Provider, db database.DbAdapter) {
	c := auth.NewRefreshTokens(db, jwt)
	group.Post("/refresh", c.Handle)
}
