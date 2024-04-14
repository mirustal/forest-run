package route

import (
	"github.com/gofiber/fiber/v2"
	"main-server/api/middleware"
	"main-server/api/route/auth"
	"main-server/database"
	"main-server/jwt"
)

func Setup(app *fiber.App, db database.DbAdapter, jwt jwt.Provider) {
	auth.InitSignUp(app, db)
	auth.InitSignIn(app, db, jwt)

	protectedRouter := app.Group("/api")
	middleware.InitAuth(protectedRouter, jwt)
	auth.InitRefreshTokens(protectedRouter, jwt, db)
}
