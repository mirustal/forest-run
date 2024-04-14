package auth

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"main-server/api/controller"
	"main-server/database"
	"main-server/domain"
	"net/http"
)

type signUp struct {
	db database.AuthRepo
}

func NewSignUp(db database.AuthRepo) controller.Controller {
	return &signUp{
		db: db,
	}
}

func (c signUp) Handle(ctx *fiber.Ctx) error {
	request := new(domain.SignUpRequest)
	if err := ctx.BodyParser(request); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(domain.ErrorResponse{Message: "can't parse request json"})
	}

	if err := Validate(*request); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(domain.ErrorResponse{Message: err.Error()})
	}

	err := c.db.StoreNewUser(request.Username, request.Password.Hash(), ctx.UserContext())
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
