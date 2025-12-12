package routes

import "github.com/faiz-muttaqin/shadcn-admin-go-starter/backend/pkg/util"

func WebSocketRoutes() {
	r := R.Group(util.GetPathOnly(util.Getenv("VITE_BACKEND", "/api")))
	r.GET("/ws")
}
