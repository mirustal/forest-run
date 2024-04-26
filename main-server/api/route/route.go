package route

import (
	"forest-run/common/defs"
	"forest-run/common/jwt"
	"forest-run/common/middleware"
	"forest-run/main-server/api/route/auth"
	"forest-run/main-server/api/route/runs"
	"forest-run/main-server/api/route/subscriptions"
	"forest-run/main-server/database"
	"forest-run/main-server/notifications"
	"forest-run/main-server/purchasing"
	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App, db database.DbAdapter, jwt jwt.Provider, notifs notifications.Manager, defs defs.Defs, purchases purchasing.Manager) {
	auth.InitSignUp(app, db)
	auth.InitSignIn(app, db, jwt)
	auth.InitRefreshTokens(app, jwt, db)

	protectedRouter := app.Group("/api")
	middleware.InitJwtAuth(protectedRouter, jwt)

	subscriptions.InitSubscribe(protectedRouter, notifs, db)
	subscriptions.InitUnsubscribe(protectedRouter, db)

	runs.InitCreate(protectedRouter, db, notifs, defs, purchases)
	runs.InitUpdate(protectedRouter, db, notifs, defs, purchases)
	runs.InitInvite(protectedRouter, db, notifs)
}
