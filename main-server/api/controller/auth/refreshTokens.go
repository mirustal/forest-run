package auth

import (
	"forest-run/main-server/api/controller"
	"forest-run/main-server/api/protocol"
	"forest-run/main-server/database"
	"forest-run/main-server/jwt"
	"github.com/gofiber/fiber/v2"
	"net/http"
	"time"
)

type refreshTokens struct {
	db  database.DbAdapter
	jwt jwt.Provider
}

func NewRefreshTokens(db database.DbAdapter, jwt jwt.Provider) controller.Controller {
	return refreshTokens{db: db, jwt: jwt}
}

func (c refreshTokens) Handle(ctx *fiber.Ctx) error {
	request := new(protocol.RefreshTokensRequest)
	if err := ctx.BodyParser(request); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(protocol.ErrorResponse{Message: "can't parse request json"})
	}

	authData, err := c.jwt.ParseUnverified(request.AuthToken)
	if err != nil {
		return ctx.Status(http.StatusUnauthorized).JSON(protocol.ErrorResponse{Message: "wrong auth token"})
	}

	activeToken, err := c.db.GetUserRefreshToken(authData.UserId, ctx.UserContext())
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(protocol.ErrorResponse{Message: "error on getting active refresh token"})
	}

	if activeToken.Token == nil || activeToken.ExpiresAt == nil {
		return ctx.Status(http.StatusUnauthorized).JSON(protocol.ErrorResponse{Message: "no refresh token is registered"})
	}

	if request.RefreshToken != *activeToken.Token {
		return ctx.Status(http.StatusUnauthorized).JSON(protocol.ErrorResponse{Message: "wrong refresh token"})
	}

	if time.Now().After(*activeToken.ExpiresAt) {
		return ctx.Status(http.StatusUnauthorized).JSON(protocol.ErrorResponse{Message: "refresh token is expired"})
	}

	refreshToken, err := c.jwt.CreateRefreshToken()
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(protocol.ErrorResponse{Message: "error on creating refresh token"})
	}

	err = c.db.StoreUserRefreshToken(authData.UserId, refreshToken, ctx.UserContext())
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(protocol.ErrorResponse{Message: "error on storing new refresh token"})
	}

	jwtToken, err := c.jwt.CreateToken(authData)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(protocol.ErrorResponse{Message: "error on creating new jwt token"})
	}

	return ctx.Status(http.StatusOK).JSON(protocol.RefreshTokensResponse{
		AuthDataResponse: protocol.AuthDataResponse{
			RefreshToken: *refreshToken.Token,
			AuthToken:    jwtToken,
		},
	})
}
