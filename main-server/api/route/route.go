package route

import (
	"github.com/gofiber/fiber/v2"
	"main-server/api/middleware"
	"main-server/database"
	"main-server/jwt"
)

func Setup(app *fiber.App, db database.DbAdapter, jwt jwt.Provider) {
	initSignUp(app, db)
	initSignIn(app, db, jwt)

	protectedRouter := app.Group("/api")
	middleware.InitAuth(protectedRouter, jwt)
	initRefreshTokens(protectedRouter, jwt, db)
}
