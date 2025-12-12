package websocket

import (
	"fmt"
	"net/http"

	"github.com/faiz-muttaqin/shadcn-admin-go-starter/backend/internal/helper"
	"github.com/gobwas/ws"
	"github.com/sirupsen/logrus"
)

func WebsocketHandler(w http.ResponseWriter, r *http.Request) {
	userData, err := helper.GetFirebaseUser2(r)
	if err != nil {
		http.Error(w, "Unauthorized: user authentication failed", http.StatusUnauthorized)
		return
	}
	fmt.Println(userData)

	// Upgrade connection
	conn, _, _, err := ws.UpgradeHTTP(r, w)
	if err != nil {
		return
	}
	if err := WS_POOL.Add(conn); err != nil {
		logrus.Printf("Failed to add connection %v", err)
		conn.Close()
	}

}
