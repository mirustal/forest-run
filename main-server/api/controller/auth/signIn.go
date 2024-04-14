package auth

import (
	"github.com/gofiber/fiber/v2"
	"main-server/api/controller"
	"main-server/database"
	"main-server/domain"
	"main-server/jwt"
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
	request := new(domain.SignInRequest)
	if err := ctx.BodyParser(request); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(domain.ErrorResponse{Message: "can't parse request json"})
	}

	if err := Validate(request.SignUpRequest); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(domain.ErrorResponse{Message: err.Error()})
	}

	user, err := s.db.GetUserByUsername(request.Username, ctx.UserContext())
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(domain.ErrorResponse{Message: "error on searching for password"})
	}

	if request.Password.Matches(user.HashedPassword) == false {
		return ctx.Status(http.StatusUnauthorized).JSON(domain.ErrorResponse{Code: domain.CodeWrongPassword, Message: "password not matches"})
	}

	rt, err := s.jwt.CreateRefreshToken()
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(domain.ErrorResponse{Message: "error on creating refresh token"})
	}

	t, err := s.jwt.CreateToken(domain.JWTBody{UserId: user.Id})
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(domain.ErrorResponse{Message: "error on creating jwt token"})
	}

	err = s.db.StoreUserRefreshToken(user.Id, rt, ctx.UserContext())
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(domain.ErrorResponse{Message: "error on storing token"})
	}

	return ctx.Status(http.StatusOK).JSON(domain.SignInResponse{
		AuthDataResponse: domain.AuthDataResponse{
			RefreshToken: *rt.Token,
			AuthToken:    t,
		},
	})
}
