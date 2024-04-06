package controller

import (
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
	"main-server/database"
	"main-server/domain"
	"net/http"
)

type SignUp struct {
	db     database.DbAdapter
	logger *zap.Logger
}

func NewSignUp(db database.DbAdapter, logger *zap.Logger) *SignUp {
	return &SignUp{
		db:     db,
		logger: logger,
	}
}

func (c *SignUp) Handle(ctx *fiber.Ctx) error {
	request := new(domain.SignUpRequest)
	if err := ctx.BodyParser(request); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(domain.ErrorResponse{Message: "can't parse request json"})
	}

	if len(request.Login) == 0 {
		return ctx.Status(http.StatusBadRequest).JSON(domain.ErrorResponse{Message: "empty login"})
	}

	if len(request.Login) > 64 {
		return ctx.Status(http.StatusBadRequest).JSON(domain.ErrorResponse{Message: "too long login"})
	}

	if len(request.Password) == 0 {
		return ctx.Status(http.StatusBadRequest).JSON(domain.ErrorResponse{Message: "empty password"})
	}

	if len(request.Password) > 64 {
		return ctx.Status(http.StatusBadRequest).JSON(domain.ErrorResponse{Message: "too long password"})
	}

	panic("not implemented")
}
