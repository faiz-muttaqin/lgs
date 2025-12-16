package audit

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Log(
	c *gin.Context,
	db *gorm.DB,
	userID uint,
	entry *Entry,
) {
	if c == nil || db == nil || entry == nil {
		return
	}

	log := LogActivity{
		UserID:    userID,
		IP:        c.ClientIP(),
		UserAgent: c.Request.UserAgent(),

		Action:     entry.Action,
		Resource:   entry.Resource,
		ResourceID: entry.ResourceID,

		ReqMethod: c.Request.Method,
		ReqURI:    c.Request.RequestURI,

		BeforeData: entry.BeforeData,
		AfterData:  entry.AfterData,

		Status:  entry.Status,
		Message: entry.Message,
	}

	// Non-blocking safety: do not break request if logging fails
	_ = db.Create(&log).Error
}
