package websockets

import (
	"net/http"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/faiz-muttaqin/lgs/backend/internal/model"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

var (
	Upgrader = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}

	// Clients maps connectionID to Client info
	Clients = make(map[string]*Client)
	// UserConnections maps userID to all their connection IDs
	UserConnections = make(map[uint][]string)
	Mutex           = sync.RWMutex{}
)

type Client struct {
	UserID       uint
	ConnectionID string
	Email        string
	Conn         *websocket.Conn
	LastActive   time.Time
}

func HandleWebSocket(w http.ResponseWriter, r *http.Request, user *model.User, db *gorm.DB) {
	conn, err := Upgrader.Upgrade(w, r, nil)
	if err != nil {
		logrus.Println(err)
		return
	}

	// Generate unique connection ID
	connectionID := uuid.New().String()

	client := &Client{
		UserID:       user.ID,
		ConnectionID: connectionID,
		Email:        user.Email.String(),
		Conn:         conn,
		LastActive:   time.Now(),
	}

	// Add client to maps
	Mutex.Lock()
	Clients[connectionID] = client
	UserConnections[user.ID] = append(UserConnections[user.ID], connectionID)
	logrus.Printf("User %s (ID: %d) connected. Connection ID: %s. Total connections for user: %d",
		user.Email, user.ID, connectionID, len(UserConnections[user.ID]))
	Mutex.Unlock()

	go HandleMessages(connectionID, client, db)
}

func HandleMessages(connectionID string, client *Client, db *gorm.DB) {
	defer func() {
		// Remove this specific connection
		Mutex.Lock()
		delete(Clients, connectionID)

		// Remove from user's connection list
		if connections, exists := UserConnections[client.UserID]; exists {
			newConnections := []string{}
			for _, connID := range connections {
				if connID != connectionID {
					newConnections = append(newConnections, connID)
				}
			}
			if len(newConnections) > 0 {
				UserConnections[client.UserID] = newConnections
			} else {
				delete(UserConnections, client.UserID)
			}
		}

		logrus.Printf("User %s (ID: %d) disconnected. Connection ID: %s. Remaining connections: %d",
			client.Email, client.UserID, connectionID, len(UserConnections[client.UserID]))
		Mutex.Unlock()

		client.Conn.Close()

		// Check if user has any remaining connections
		go checkForReconnection(client.UserID, client.Email, db)
	}()

	for {
		messageType, p, err := client.Conn.ReadMessage()
		if err != nil {
			logrus.Println(err)
			return
		}

		// Update last active time
		Mutex.Lock()
		client.LastActive = time.Now()
		Mutex.Unlock()

		logrus.Printf("Received message from user %s (connection %s): %s", client.Email, connectionID, string(p))

		// Handle the message
		HandleMessage(messageType, p, client)
	}
}

func HandleMessage(messageType int, message []byte, sender *Client) {
	// Example: Assume messages have the format "recipientUserID:message"
	parts := strings.SplitN(string(message), ":", 2)
	if len(parts) != 2 {
		logrus.Println("Invalid message format:", string(message))
		return
	}

	recipientUserID, err := strconv.ParseUint(parts[0], 10, 32)
	if err != nil {
		logrus.Println("Invalid recipient user ID:", parts[0])
		return
	}

	actualMessage := parts[1]

	// Send to specific user (all their tabs)
	SendMessageToUser(messageType, actualMessage, uint(recipientUserID))
}

// SendMessageToUser sends message to ALL tabs/connections of a specific user
func SendMessageToUser(messageType int, message string, userID uint) {
	Mutex.RLock()
	connectionIDs := UserConnections[userID]
	Mutex.RUnlock()

	if len(connectionIDs) == 0 {
		logrus.Printf("User ID %d is not connected", userID)
		return
	}

	successCount := 0
	Mutex.RLock()
	defer Mutex.RUnlock()

	for _, connID := range connectionIDs {
		if client, ok := Clients[connID]; ok && client.Conn != nil {
			err := client.Conn.WriteMessage(messageType, []byte(message))
			if err != nil {
				logrus.Printf("Error sending to connection %s: %v", connID, err)
			} else {
				successCount++
			}
		}
	}

	logrus.Printf("Sent message to user %d (%d/%d connections)", userID, successCount, len(connectionIDs))
}

// SendMessageToConnection sends message to a specific connection/tab
func SendMessageToConnection(messageType int, message string, connectionID string) {
	Mutex.RLock()
	defer Mutex.RUnlock()

	if client, ok := Clients[connectionID]; ok && client.Conn != nil {
		err := client.Conn.WriteMessage(messageType, []byte(message))
		if err != nil {
			logrus.Println(err)
		}
	}
}

// CloseUserConnections closes all connections for a user
func CloseUserConnections(userID uint) {
	Mutex.Lock()
	connectionIDs := UserConnections[userID]
	Mutex.Unlock()

	for _, connID := range connectionIDs {
		Mutex.RLock()
		if client, ok := Clients[connID]; ok && client.Conn != nil {
			client.Conn.Close()
		}
		Mutex.RUnlock()
	}

	Mutex.Lock()
	for _, connID := range connectionIDs {
		delete(Clients, connID)
	}
	delete(UserConnections, userID)
	Mutex.Unlock()

	logrus.Printf("Closed all connections for user ID %d", userID)
}

func checkForReconnection(userID uint, email string, db *gorm.DB) {
	disconectionTimeStr := os.Getenv("MAX_DISCONECTION_TIME_S")
	disconectionExpiredSeconds, err := strconv.Atoi(disconectionTimeStr)
	if err != nil {
		disconectionExpiredSeconds = 30
	}

	// Wait for timeout
	time.Sleep(time.Duration(disconectionExpiredSeconds) * time.Second)

	// Check if user still has ANY active connections
	Mutex.RLock()
	hasConnections := len(UserConnections[userID]) > 0
	Mutex.RUnlock()

	if !hasConnections {
		logrus.Printf("User %s (ID: %d) has no active connections after %d seconds", email, userID, disconectionExpiredSeconds)
		updates := map[string]interface{}{
			"Session":        "",
			"SessionExpired": 0,
		}

		if err := db.Model(&model.User{}).Where("id = ?", userID).Updates(updates).Error; err != nil {
			logrus.Errorf("Failed to update user session: %v", err)
			return
		}
	}
}

// IsUserConnected checks if a user has any active connections
func IsUserConnected(userID uint) bool {
	Mutex.RLock()
	defer Mutex.RUnlock()

	return len(UserConnections[userID]) > 0
}

// GetUserConnectionCount returns number of active connections for a user
func GetUserConnectionCount(userID uint) int {
	Mutex.RLock()
	defer Mutex.RUnlock()

	return len(UserConnections[userID])
}

// BroadcastMessage sends a message to all connected clients (all users, all tabs)
func BroadcastMessage(messageType int, message string) {
	Mutex.RLock()
	defer Mutex.RUnlock()

	count := 0
	for _, client := range Clients {
		if client.Conn != nil {
			err := client.Conn.WriteMessage(messageType, []byte(message))
			if err != nil {
				logrus.Errorf("Error broadcasting to connection %s: %v", client.ConnectionID, err)
			} else {
				count++
			}
		}
	}

	logrus.Printf("Broadcast message sent to %d connections", count)
}

// BroadcastToUsers sends message to specific list of users (all their tabs)
func BroadcastToUsers(messageType int, message string, userIDs []uint) {
	for _, userID := range userIDs {
		SendMessageToUser(messageType, message, userID)
	}
}

// GetOnlineUsers returns list of currently connected user IDs
func GetOnlineUsers() []uint {
	Mutex.RLock()
	defer Mutex.RUnlock()

	userIDs := make([]uint, 0, len(UserConnections))
	for userID := range UserConnections {
		userIDs = append(userIDs, userID)
	}

	return userIDs
}

// GetConnectionStats returns connection statistics
func GetConnectionStats() map[string]interface{} {
	Mutex.RLock()
	defer Mutex.RUnlock()

	return map[string]interface{}{
		"total_connections": len(Clients),
		"total_users":       len(UserConnections),
		"online_users":      GetOnlineUsers(),
	}
}
