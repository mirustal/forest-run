package controller

import "github.com/gofiber/fiber/v2"

type connect struct {
}

func NewConnect() Controller {
	return &connect{}
}

func (c connect) Handle(ctx *fiber.Ctx) error {
	// todo
	return nil
}
