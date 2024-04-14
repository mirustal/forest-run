package route

import (
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
	"main-server/boot"
	"main-server/database"
	"main-server/jwt"
)

func Setup(app *fiber.App, env boot.Env, logger *zap.Logger, db database.DbAdapter, jwt jwt.Provider) {
	defer logger.Sync()
	initSignUp(app, db, logger)
	initSignIn(app, db, logger, jwt)
}
