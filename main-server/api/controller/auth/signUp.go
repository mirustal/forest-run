package auth

import (
	"errors"
	"forest-run/main-server/api/controller"
	"forest-run/main-server/api/protocol"
	"forest-run/main-server/database"
	"github.com/gofiber/fiber/v2"
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
	request := new(protocol.SignUpRequest)
	if err := ctx.BodyParser(request); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(protocol.ErrorResponse{Message: "can't parse request json"})
	}

	if err := Validate(*request); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(protocol.ErrorResponse{Message: err.Error()})
	}

	err := c.db.StoreNewUser(request.Username, hash(request.Password), ctx.UserContext())
	if err != nil {
		var usernameAlreadyTakenError database.UsernameAlreadyTakenError
		if errors.As(err, &usernameAlreadyTakenError) {
			return ctx.Status(http.StatusConflict).JSON(protocol.ErrorResponse{
				Code:    protocol.CodeUserNameAlreadyTaken,
				Message: "username already taken",
			})
		}
		return ctx.Status(http.StatusInternalServerError).JSON(protocol.ErrorResponse{Message: "unable to store user"})
	}

	return ctx.Status(http.StatusOK).JSON(protocol.SignUpResponse{})
}
