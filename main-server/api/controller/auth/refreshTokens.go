package auth

import (
	"github.com/gofiber/fiber/v2"
	"main-server/api/controller"
	"main-server/database"
	"main-server/domain"
	"main-server/jwt"
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
	request := new(domain.RefreshTokensRequest)
	if err := ctx.BodyParser(request); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(domain.ErrorResponse{Message: "can't parse request json"})
	}

	authData, err := c.jwt.ParseUnverified(request.AuthToken)
	if err != nil {
		return ctx.Status(http.StatusUnauthorized).JSON(domain.ErrorResponse{Message: "wrong auth token"})
	}

	activeToken, err := c.db.GetUserRefreshToken(authData.UserId, ctx.UserContext())
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(domain.ErrorResponse{Message: "error on getting active refresh token"})
	}

	if activeToken.Token == nil || activeToken.ExpiresAt == nil {
		return ctx.Status(http.StatusUnauthorized).JSON(domain.ErrorResponse{Message: "no refresh token is registered"})
	}

	if request.RefreshToken != *activeToken.Token {
		return ctx.Status(http.StatusUnauthorized).JSON(domain.ErrorResponse{Message: "wrong refresh token"})
	}

	if time.Now().After(*activeToken.ExpiresAt) {
		return ctx.Status(http.StatusUnauthorized).JSON(domain.ErrorResponse{Message: "refresh token is expired"})
	}

	refreshToken, err := c.jwt.CreateRefreshToken()
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(domain.ErrorResponse{Message: "error on creating refresh token"})
	}

	err = c.db.StoreUserRefreshToken(authData.UserId, refreshToken, ctx.UserContext())
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(domain.ErrorResponse{Message: "error on storing new refresh token"})
	}

	jwtToken, err := c.jwt.CreateToken(authData)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(domain.ErrorResponse{Message: "error on creating new jwt token"})
	}

	return ctx.Status(http.StatusOK).JSON(domain.RefreshTokensResponse{
		AuthDataResponse: domain.AuthDataResponse{
			RefreshToken: *refreshToken.Token,
			AuthToken:    jwtToken,
		},
	})
}
