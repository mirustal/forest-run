package controller

import "github.com/gofiber/fiber/v2"

type Controller interface {
	Handle(ctx *fiber.Ctx) error
}
