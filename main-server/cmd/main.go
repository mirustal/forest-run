package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"main/api/route"
	"main/boot"
)

func main() {
	app := fiber.New()
	env := boot.NewEnv()
	// db := boot.NewDb()

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept, Authorization",
	}))

	route.Setup(app, env /*, db*/)

	if err := app.Listen(":1337"); err != nil {
		fmt.Printf("Oops... Server is not running! Reason: %v", err)
	}
}
