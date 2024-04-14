package middleware

import (
	"github.com/gofiber/fiber/v2"
	"main-server/domain"
	"main-server/jwt"
	"net/http"
	"strings"
)

const authDataKey = "jwt.auth-data"

func InitAuth(group fiber.Router, jwt jwt.Provider) {
	group.Use(auth{
		jwt: jwt,
	}.authorize)
}

type auth struct {
	jwt jwt.Provider
}

func (a auth) authorize(ctx *fiber.Ctx) error {
	token := ctx.Get("Authorization", "")

	if strings.HasPrefix(token, "Bearer ") == false {
		return ctx.Status(http.StatusBadRequest).JSON(domain.ErrorResponse{
			Message: "authorization header must start with Bearer",
		})
	}

	token = strings.Replace(token, "Bearer ", "", 1)

	if token == "" {
		return ctx.Status(http.StatusBadRequest).JSON(domain.ErrorResponse{
			Message: "authorization header not provided",
		})
	}

	authData, err := a.jwt.Parse(domain.JWTToken(token))
	if err != nil {
		ctx.Context().Logger().Printf("error while parsing jwt token: %v", err.Error())
		return ctx.Status(http.StatusUnauthorized).JSON(domain.ErrorResponse{
			Message: "error while parsing jwt token",
		})
	}

	ctx.Locals(authDataKey, authData)
	return ctx.Next()
}

func GetAuthData(c *fiber.Ctx) domain.JWTBody {
	return c.Locals(authDataKey).(domain.JWTBody)
}
