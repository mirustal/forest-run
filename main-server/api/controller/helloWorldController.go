package controller

import (
	"github.com/gofiber/fiber/v2"
	"main-server/domain"
)

type HelloWorldController struct {
}

func NewHelloWorldController() HelloWorldController {
	return HelloWorldController{}
}

func (c HelloWorldController) Handle(ctx *fiber.Ctx) error {
	return ctx.Status(200).JSON(domain.HelloWorldResponse{Message: "Hello World"})
}
