package controller

import (
	"github.com/gofiber/contrib/websocket"
)

type WsController interface {
	Handle(ctx *websocket.Conn)
}
