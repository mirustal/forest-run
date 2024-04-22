package route

import (
	"github.com/gofiber/fiber/v2"
	"main-server/api/middleware"
	"main-server/api/route/auth"
	"main-server/api/route/runs"
	"main-server/api/route/subscriptions"
	"main-server/database"
	"main-server/defs"
	"main-server/jwt"
	"main-server/notifications"
	"main-server/purchasing"
)

func Setup(app *fiber.App, db database.DbAdapter, jwt jwt.Provider, notifs notifications.Manager, defs defs.Defs, purchases purchasing.Manager) {
	auth.InitSignUp(app, db)
	auth.InitSignIn(app, db, jwt)
	auth.InitRefreshTokens(app, jwt, db)

	protectedRouter := app.Group("/api")
	middleware.InitAuth(protectedRouter, jwt)

	subscriptions.InitSubscribe(protectedRouter, notifs, db)
	subscriptions.InitUnsubscribe(protectedRouter, db)

	runs.InitCreate(protectedRouter, db, notifs, defs, purchases)
	runs.InitUpdate(protectedRouter, db, notifs, defs, purchases)
	runs.InitInvite(protectedRouter, db, notifs)
}
