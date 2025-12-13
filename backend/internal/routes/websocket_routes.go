package routes

import (
	"github.com/faiz-muttaqin/lgs/backend/internal/websockets"
	"github.com/faiz-muttaqin/lgs/backend/pkg/util"
)

func WebSocketRoutes() {
	r := R.Group(util.GetPathOnly(util.Getenv("VITE_BACKEND", "/api")))
	r.GET("/ws", websockets.WebsocketHandlerGin)
}
