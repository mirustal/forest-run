package route

import (
	"github.com/gofiber/fiber/v2"
	"main-server/database"
	"main-server/jwt"
)

//	@Summary		Refresh Tokens
//	@Description	hello bitch
//	@Tags			auth
//	@Produce		json
//	@Security		ApiKeyAuth
//	@Param			input	body		domain.RefreshTokensRequest	true	"Refresh tokens data"
//	@Success		200		{object}	domain.RefreshTokensResponse
//	@Failure		400,401	{object}	domain.ErrorResponse
//	@Failure		500		{object}	domain.ErrorResponse
//	@Router			/api/refresh [get]
func initRefreshTokens(fiber.Router, jwt.Provider, database.DbAdapter) {

}
