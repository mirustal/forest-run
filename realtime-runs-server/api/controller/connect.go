package controller

import (
	"forest-run/common/defs"
	"forest-run/common/middleware"
	"github.com/gofiber/contrib/websocket"
	"time"
)

type connect struct {
	defs defs.Defs
}

func NewConnect(defs defs.Defs) WsController {
	return &connect{defs: defs}
}

func (c connect) Handle(conn *websocket.Conn) {
	authData := middleware.GetAuthDataWs(conn)

	conn.WriteJSON(authData)
	go TDelayed(conn)
}

func TDelayed(conn *websocket.Conn) {
	time.Sleep(time.Second * 10)
	conn.WriteJSON(struct {
		Message string
	}{
		Message: "delayed message",
	})
	conn.Close()
}
