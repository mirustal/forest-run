package controller

import (
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
	"main-server/database"
	"main-server/domain"
	"net/http"
	"time"
)

type signIn struct {
	db     database.DbAdapter
	logger *zap.Logger
}

func NewSignIn(db database.DbAdapter, logger *zap.Logger) Controller {
	return &signIn{db: db, logger: logger}
}

func (s signIn) Handle(ctx *fiber.Ctx) error {
	request := new(domain.SignInRequest)
	if err := ctx.BodyParser(request); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(domain.ErrorResponse{Message: "can't parse request json"})
	}

	if err := request.Validate(); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(domain.ErrorResponse{Message: err.Error()})
	}

	user, err := s.db.GetUserByUsername(request.Username, ctx.UserContext())
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(domain.ErrorResponse{Message: "error on searching for password"})
	}

	if request.Password.Matches(user.HashedPassword) == false {
		return ctx.Status(http.StatusUnauthorized).JSON(domain.ErrorResponse{Code: domain.CodeWrongPassword, Message: "password not matches"})
	}

	rt := domain.NewRefreshToken("TODO")
	jwt := domain.NewJWTToken(domain.JWTBody{ /*TODO*/ }, "TODO")

	rtData := domain.RefreshTokenData{
		Token:     rt,
		ExpiresAt: time.Time{}, // TODO
	}

	err = s.db.StoreUserRefreshToken(user.Id, rtData, ctx.UserContext())

	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(domain.ErrorResponse{Message: "error on storing token"})
	}

	return ctx.Status(http.StatusOK).JSON(domain.SignInResponse{
		AuthDataResponse: domain.AuthDataResponse{
			RefreshToken: rtData,
			AuthToken:    jwt,
		},
	})
}
