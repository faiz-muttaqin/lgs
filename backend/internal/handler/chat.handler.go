package handler

import (
	"net/http"
	"strconv"
	"time"

	"github.com/faiz-muttaqin/lgs/backend/internal/helper"
	"github.com/faiz-muttaqin/lgs/backend/internal/model"
	"github.com/faiz-muttaqin/lgs/backend/pkg/audit"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ChatHandler struct {
	DB *gorm.DB
}

func NewChatHandler(db *gorm.DB) *ChatHandler {
	return &ChatHandler{DB: db}
}

// GetMyChats retrieves all chats for the authenticated user
func GetMyChats(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {

		userData, err := helper.GetFirebaseUser(c)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"success": false,
				"message": "Unauthorized",
			})
			return
		}

		var chats []model.Chat
		if err := db.Where("user1_id = ? OR user2_id = ?", userData.ID, userData.ID).
			Preload("User1").
			Preload("User2").
			Preload("Product").
			Order("updated_at DESC").
			Find(&chats).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"success": false,
				"message": "Failed to retrieve chats",
			})
			return
		}

		// Populate virtual fields for each chat
		for i := range chats {
			// Get last message
			var lastMsg model.Message
			if err := db.Where("chat_id = ?", chats[i].ID).
				Order("created_at DESC").
				First(&lastMsg).Error; err == nil {
				chats[i].LastMessage = &lastMsg
			}

			// Count unread messages
			var unreadCount int64
			db.Model(&model.Message{}).
				Where("chat_id = ? AND receiver_id = ? AND read_at IS NULL", chats[i].ID, userData.ID).
				Count(&unreadCount)
			chats[i].UnreadCount = int(unreadCount)

			// Set other user
			if chats[i].User1ID == userData.ID {
				chats[i].OtherUser = &chats[i].User2
			} else {
				chats[i].OtherUser = &chats[i].User1
			}
		}

		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"data": gin.H{
				"chats": chats,
				"count": len(chats),
			},
		})
	}
}

// GetOrCreateChat gets existing chat or creates new one between two users
func GetOrCreateChat(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {

		userData, err := helper.GetFirebaseUser(c)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"success": false,
				"message": "Unauthorized",
			})
			return
		}

		var input struct {
			OtherUserID uint  `json:"other_user_id" binding:"required"`
			ProductID   *uint `json:"product_id"`
		}

		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"success": false,
				"message": "Invalid input",
			})
			return
		}

		// Don't allow chat with self
		if input.OtherUserID == userData.ID {
			c.JSON(http.StatusBadRequest, gin.H{
				"success": false,
				"message": "Cannot create chat with yourself",
			})
			return
		}

		// Check if chat already exists
		var chat model.Chat
		err = db.Where("(user1_id = ? AND user2_id = ?) OR (user1_id = ? AND user2_id = ?)",
			userData.ID, input.OtherUserID, input.OtherUserID, userData.ID).
			First(&chat).Error

		if err == gorm.ErrRecordNotFound {
			// Create new chat
			chat = model.Chat{
				User1ID: userData.ID,
				User2ID: input.OtherUserID,
			}
			if input.ProductID != nil {
				chat.ProductID.Int64 = int64(*input.ProductID)
				chat.ProductID.Valid = true
			}

			if err := db.Create(&chat).Error; err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{
					"success": false,
					"message": "Failed to create chat",
				})
				return
			}

			audit.Log(c, db, userData, audit.Create("chat", chat.ID).After(chat).Success("Chat created"))
		}

		// Load relations
		db.Preload("User1").Preload("User2").Preload("Product").First(&chat, chat.ID)

		// Set other user
		if chat.User1ID == userData.ID {
			chat.OtherUser = &chat.User2
		} else {
			chat.OtherUser = &chat.User1
		}

		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"data":    chat,
		})
	}

}

// GetChatMessages retrieves messages for a specific chat
func GetChatMessages(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {

		userData, err := helper.GetFirebaseUser(c)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"success": false,
				"message": "Unauthorized",
			})
			return
		}

		chatID, _ := strconv.ParseUint(c.Param("id"), 10, 32)

		// Verify user is part of the chat
		var chat model.Chat
		if err := db.Where("id = ? AND (user1_id = ? OR user2_id = ?)", chatID, userData.ID, userData.ID).
			First(&chat).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{
				"success": false,
				"message": "Chat not found",
			})
			return
		}

		// Get pagination params
		page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
		limit, _ := strconv.Atoi(c.DefaultQuery("limit", "50"))
		offset := (page - 1) * limit

		var messages []model.Message
		if err := db.Where("chat_id = ? AND is_deleted = false", chatID).
			Preload("Sender").
			Preload("Receiver").
			Preload("ReplyToMessage").
			Preload("ReplyToMessage.Sender").
			Order("created_at DESC").
			Limit(limit).
			Offset(offset).
			Find(&messages).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"success": false,
				"message": "Failed to retrieve messages",
			})
			return
		}

		// Count total messages
		var total int64
		db.Model(&model.Message{}).Where("chat_id = ? AND is_deleted = false", chatID).Count(&total)

		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"data": gin.H{
				"messages": messages,
				"total":    total,
				"page":     page,
				"limit":    limit,
			},
		})
	}

}

// SendMessage sends a new message in a chat
func SendMessage(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {

		userData, err := helper.GetFirebaseUser(c)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"success": false,
				"message": "Unauthorized",
			})
			return
		}

		chatID, _ := strconv.ParseUint(c.Param("id"), 10, 32)

		// Verify user is part of the chat
		var chat model.Chat
		if err := db.Where("id = ? AND (user1_id = ? OR user2_id = ?)", chatID, userData.ID, userData.ID).
			First(&chat).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{
				"success": false,
				"message": "Chat not found",
			})
			return
		}

		var input struct {
			Content          string `json:"content" binding:"required"`
			AttachmentURL    string `json:"attachment_url"`
			AttachmentType   string `json:"attachment_type"`
			ReplyToMessageID *uint  `json:"reply_to_message_id"`
		}

		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"success": false,
				"message": "Invalid input",
			})
			return
		}

		// Determine receiver
		var receiverID uint
		if chat.User1ID == userData.ID {
			receiverID = chat.User2ID
		} else {
			receiverID = chat.User1ID
		}

		message := model.Message{
			ChatID:         uint(chatID),
			SenderID:       userData.ID,
			ReceiverID:     receiverID,
			Content:        input.Content,
			AttachmentURL:  input.AttachmentURL,
			AttachmentType: input.AttachmentType,
		}

		if input.ReplyToMessageID != nil {
			message.ReplyToMessageID.Int64 = int64(*input.ReplyToMessageID)
			message.ReplyToMessageID.Valid = true
		}

		if err := db.Create(&message).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"success": false,
				"message": "Failed to send message",
			})
			return
		}

		// Update chat's updated_at
		db.Model(&chat).Update("updated_at", time.Now())

		// Load relations
		db.Preload("Sender").
			Preload("Receiver").
			Preload("ReplyToMessage").
			Preload("ReplyToMessage.Sender").
			First(&message, message.ID)

		audit.Log(c, db, userData, audit.Create("message", message.ID).After(message).Success("Message sent"))

		c.JSON(http.StatusCreated, gin.H{
			"success": true,
			"message": "Message sent",
			"data":    message,
		})
	}

}

// MarkMessageReceived marks a message as received
func MarkMessageReceived(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {

		userData, err := helper.GetFirebaseUser(c)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"success": false,
				"message": "Unauthorized",
			})
			return
		}

		messageID, _ := strconv.ParseUint(c.Param("id"), 10, 32)

		var message model.Message
		if err := db.Where("id = ? AND receiver_id = ?", messageID, userData.ID).First(&message).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{
				"success": false,
				"message": "Message not found",
			})
			return
		}

		if message.ReceivedAt == nil {
			now := time.Now()
			message.ReceivedAt = &now
			db.Save(&message)
		}

		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"message": "Message marked as received",
		})
	}

}

// MarkMessageRead marks a message as read
func MarkMessageRead(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {

		userData, err := helper.GetFirebaseUser(c)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"success": false,
				"message": "Unauthorized",
			})
			return
		}

		messageID, _ := strconv.ParseUint(c.Param("id"), 10, 32)

		var message model.Message
		if err := db.Where("id = ? AND receiver_id = ?", messageID, userData.ID).First(&message).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{
				"success": false,
				"message": "Message not found",
			})
			return
		}

		if message.ReadAt == nil {
			now := time.Now()
			message.ReadAt = &now
			if message.ReceivedAt == nil {
				message.ReceivedAt = &now
			}
			db.Save(&message)
		}

		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"message": "Message marked as read",
		})
	}

}

// MarkChatRead marks all messages in a chat as read
func MarkChatRead(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {

		userData, err := helper.GetFirebaseUser(c)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"success": false,
				"message": "Unauthorized",
			})
			return
		}

		chatID, _ := strconv.ParseUint(c.Param("id"), 10, 32)

		// Verify user is part of the chat
		var chat model.Chat
		if err := db.Where("id = ? AND (user1_id = ? OR user2_id = ?)", chatID, userData.ID, userData.ID).
			First(&chat).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{
				"success": false,
				"message": "Chat not found",
			})
			return
		}

		now := time.Now()
		db.Model(&model.Message{}).
			Where("chat_id = ? AND receiver_id = ? AND read_at IS NULL", chatID, userData.ID).
			Updates(map[string]interface{}{
				"read_at":     now,
				"received_at": now,
			})

		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"message": "All messages marked as read",
		})
	}

}

// EditMessage edits a message (within 7 minutes and not read)
func EditMessage(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {

		userData, err := helper.GetFirebaseUser(c)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"success": false,
				"message": "Unauthorized",
			})
			return
		}

		messageID, _ := strconv.ParseUint(c.Param("id"), 10, 32)

		var message model.Message
		if err := db.Where("id = ? AND sender_id = ?", messageID, userData.ID).First(&message).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{
				"success": false,
				"message": "Message not found",
			})
			return
		}

		// Check if can edit
		timeSinceSent := time.Since(message.CreatedAt)
		if timeSinceSent > 7*time.Minute {
			c.JSON(http.StatusBadRequest, gin.H{
				"success": false,
				"message": "Cannot edit message after 7 minutes",
			})
			return
		}

		if message.ReadAt != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"success": false,
				"message": "Cannot edit message that has been read",
			})
			return
		}

		var input struct {
			Content string `json:"content" binding:"required"`
		}

		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"success": false,
				"message": "Invalid input",
			})
			return
		}

		oldMessage := message
		message.Content = input.Content
		message.IsEdited = true

		if err := db.Save(&message).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"success": false,
				"message": "Failed to edit message",
			})
			return
		}

		audit.Log(c, db, userData, audit.Update("message", message.ID).Before(oldMessage).After(message).Success("Message edited"))

		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"message": "Message edited",
			"data":    message,
		})
	}
}

// DeleteMessage deletes a message (within 7 minutes and not read)
func DeleteMessage(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {

		userData, err := helper.GetFirebaseUser(c)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"success": false,
				"message": "Unauthorized",
			})
			return
		}

		messageID, _ := strconv.ParseUint(c.Param("id"), 10, 32)

		var message model.Message
		if err := db.Where("id = ? AND sender_id = ?", messageID, userData.ID).First(&message).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{
				"success": false,
				"message": "Message not found",
			})
			return
		}

		// Check if can delete
		timeSinceSent := time.Since(message.CreatedAt)
		if timeSinceSent > 7*time.Minute {
			c.JSON(http.StatusBadRequest, gin.H{
				"success": false,
				"message": "Cannot delete message after 7 minutes",
			})
			return
		}

		if message.ReadAt != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"success": false,
				"message": "Cannot delete message that has been read",
			})
			return
		}

		message.IsDeleted = true
		message.Content = "[Message deleted]"
		db.Save(&message)

		audit.Log(c, db, userData, audit.Delete("message", message.ID).Before(message).Success("Message deleted"))

		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"message": "Message deleted",
		})

	}
}

// GetUnreadCount gets total unread message count for user
func GetUnreadCount(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {

		userData, err := helper.GetFirebaseUser(c)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"success": false,
				"message": "Unauthorized",
			})
			return
		}

		var count int64
		db.Model(&model.Message{}).
			Where("receiver_id = ? AND read_at IS NULL", userData.ID).
			Count(&count)

		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"data": gin.H{
				"unread_count": count,
			},
		})

	}
}

// GetUnreadMessages gets all unread messages for user
func GetUnreadMessages(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {

		userData, err := helper.GetFirebaseUser(c)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"success": false,
				"message": "Unauthorized",
			})
			return
		}

		var messages []model.Message
		if err := db.Where("receiver_id = ? AND read_at IS NULL AND is_deleted = false", userData.ID).
			Preload("Sender").
			Preload("Chat").
			Order("created_at DESC").
			Find(&messages).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"success": false,
				"message": "Failed to retrieve unread messages",
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"data": gin.H{
				"messages": messages,
				"count":    len(messages),
			},
		})
	}

}
