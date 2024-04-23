package auth

import (
	"forest-run/main-server/api/controller"
	"forest-run/main-server/api/protocol"
	"forest-run/main-server/database"
	"forest-run/main-server/domain"
	"forest-run/main-server/jwt"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

type signIn struct {
	db  database.DbAdapter
	jwt jwt.Provider
}

func NewSignIn(db database.DbAdapter, jwt jwt.Provider) controller.Controller {
	return &signIn{db: db, jwt: jwt}
}

func (s signIn) Handle(ctx *fiber.Ctx) error {
	request := new(protocol.SignInRequest)
	if err := ctx.BodyParser(request); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(protocol.ErrorResponse{Message: "can't parse request json"})
	}

	if err := Validate(request.SignUpRequest); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(protocol.ErrorResponse{Message: err.Error()})
	}

	user, err := s.db.GetUserByUsername(request.Username, ctx.UserContext())
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(protocol.ErrorResponse{Message: "error on searching for password"})
	}

	if matches(request.Password, user.HashedPassword) == false {
		return ctx.Status(http.StatusUnauthorized).JSON(protocol.ErrorResponse{Code: protocol.CodeWrongPassword, Message: "password not matches"})
	}

	rt, err := s.jwt.CreateRefreshToken()
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(protocol.ErrorResponse{Message: "error on creating refresh token"})
	}

	t, err := s.jwt.CreateToken(domain.JWTBody{UserId: user.Id})
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(protocol.ErrorResponse{Message: "error on creating jwt token"})
	}

	err = s.db.StoreUserRefreshToken(user.Id, rt, ctx.UserContext())
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(protocol.ErrorResponse{Message: "error on storing token"})
	}

	return ctx.Status(http.StatusOK).JSON(protocol.SignInResponse{
		AuthDataResponse: protocol.AuthDataResponse{
			RefreshToken: *rt.Token,
			AuthToken:    t,
		},
	})
}
