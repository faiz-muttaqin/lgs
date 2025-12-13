package websockets

import (
	"net/http"

	"github.com/faiz-muttaqin/lgs/backend/internal/database"
	"github.com/faiz-muttaqin/lgs/backend/internal/helper"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func WebsocketHandlerGin(c *gin.Context) {
	userData, err := helper.GetFirebaseUser(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Unauthorized"})
		return
	}

	logrus.Printf("WebSocket connection attempt from user: %s (ID: %d)", userData.Email, userData.ID)
	HandleWebSocket(c.Writer, c.Request, userData, database.DB)
}

// func WebsocketHandler(w http.ResponseWriter, r *http.Request) {
// 	userData, err := helper.GetFirebaseUser2(r)
// 	if err != nil {
// 		http.Error(w, "Unauthorized: user authentication failed", http.StatusUnauthorized)
// 		return
// 	}
// 	fmt.Println(userData)

// 	// Upgrade connection
// 	conn, _, _, err := ws.UpgradeHTTP(r, w)
// 	if err != nil {
// 		return
// 	}
// 	if err := WS_POOL.Add(conn); err != nil {
// 		logrus.Printf("Failed to add connection %v", err)
// 		conn.Close()
// 	}

// }
