package controller

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
	"main-server/database"
	"main-server/domain"
	"net/http"
)

type SignUp struct {
	db     database.AuthRepo
	logger *zap.Logger
}

func NewSignUp(db database.AuthRepo, logger *zap.Logger) *SignUp {
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

	if len(request.Username) == 0 {
		return ctx.Status(http.StatusBadRequest).JSON(domain.ErrorResponse{Message: "empty login"})
	}

	if len(request.Username) > 64 {
		return ctx.Status(http.StatusBadRequest).JSON(domain.ErrorResponse{Message: "too long login"})
	}

	if len(request.Password) == 0 {
		return ctx.Status(http.StatusBadRequest).JSON(domain.ErrorResponse{Message: "empty password"})
	}

	if len(request.Password) > 64 {
		return ctx.Status(http.StatusBadRequest).JSON(domain.ErrorResponse{Message: "too long password"})
	}

	err := c.db.StoreNewUser(request.Username, request.Password, ctx.UserContext())
	if err != nil {
		var usernameAlreadyTakenError database.UsernameAlreadyTakenError
		if errors.As(err, &usernameAlreadyTakenError) {
			return ctx.Status(http.StatusConflict).JSON(domain.ErrorResponse{
				Code:    domain.CodeUserNameAlreadyTaken,
				Message: "username already taken",
			})
		}
		return ctx.Status(http.StatusInternalServerError).JSON(domain.ErrorResponse{Message: "unable to store user"})
	}

	return ctx.Status(http.StatusOK).JSON(domain.SignUpResponse{})

}
