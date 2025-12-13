package model

import (
	"database/sql"
	"time"

	"gorm.io/gorm"
)

// Chat represents a one-to-one conversation between two users
type Chat struct {
	ID        uint           `gorm:"primaryKey;column:id" json:"id"`
	User1ID   uint           `gorm:"column:user1_id;not null;index:idx_chat_users" json:"user1_id"`
	User2ID   uint           `gorm:"column:user2_id;not null;index:idx_chat_users" json:"user2_id"`
	ProductID sql.NullInt64  `gorm:"column:product_id;index" json:"product_id,omitempty"` // Optional: chat started from a product
	CreatedAt time.Time      `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time      `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`

	// Relations
	User1    User      `gorm:"foreignKey:User1ID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"user1,omitempty"`
	User2    User      `gorm:"foreignKey:User2ID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"user2,omitempty"`
	Product  *Product  `gorm:"foreignKey:ProductID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL" json:"product,omitempty"`
	Messages []Message `gorm:"foreignKey:ChatID" json:"messages,omitempty"`

	// Virtual fields (not in DB, populated by queries)
	LastMessage *Message `gorm:"-" json:"last_message,omitempty"`
	UnreadCount int      `gorm:"-" json:"unread_count"`
	OtherUser   *User    `gorm:"-" json:"other_user,omitempty"` // The other participant (not current user)
}

func (Chat) TableName() string {
	return "chats"
}

// Message represents a single message in a chat
type Message struct {
	ID               uint           `gorm:"primaryKey;column:id" json:"id"`
	ChatID           uint           `gorm:"column:chat_id;not null;index" json:"chat_id"`
	SenderID         uint           `gorm:"column:sender_id;not null;index" json:"sender_id"`
	ReceiverID       uint           `gorm:"column:receiver_id;not null;index" json:"receiver_id"`
	Content          string         `gorm:"column:content;type:text;not null" json:"content"`
	AttachmentURL    string         `gorm:"column:attachment_url;size:500" json:"attachment_url,omitempty"`
	AttachmentType   string         `gorm:"column:attachment_type;size:50" json:"attachment_type,omitempty"` // image, file, etc
	ReplyToMessageID sql.NullInt64  `gorm:"column:reply_to_message_id;index" json:"reply_to_message_id,omitempty"`
	IsEdited         bool           `gorm:"column:is_edited;default:false" json:"is_edited"`
	IsDeleted        bool           `gorm:"column:is_deleted;default:false" json:"is_deleted"`
	ReceivedAt       *time.Time     `gorm:"column:received_at" json:"received_at,omitempty"`
	ReadAt           *time.Time     `gorm:"column:read_at" json:"read_at,omitempty"`
	CreatedAt        time.Time      `gorm:"column:created_at" json:"created_at"`
	UpdatedAt        time.Time      `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt        gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`

	// Relations
	Chat           Chat     `gorm:"foreignKey:ChatID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"chat,omitempty"`
	Sender         User     `gorm:"foreignKey:SenderID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"sender,omitempty"`
	Receiver       User     `gorm:"foreignKey:ReceiverID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"receiver,omitempty"`
	ReplyToMessage *Message `gorm:"foreignKey:ReplyToMessageID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL" json:"reply_to,omitempty"`

	// Virtual fields
	CanEdit   bool `gorm:"-" json:"can_edit"`   // Can be edited (within 7 mins and not read)
	CanDelete bool `gorm:"-" json:"can_delete"` // Can be deleted (within 7 mins and not read)
}

func (Message) TableName() string {
	return "messages"
}

// AfterFind hook to populate virtual fields
func (m *Message) AfterFind(tx *gorm.DB) error {
	// Calculate if message can be edited/deleted
	// Rules: within 7 minutes AND not read yet
	timeSinceSent := time.Since(m.CreatedAt)
	withinTimeLimit := timeSinceSent <= 7*time.Minute
	notRead := m.ReadAt == nil

	m.CanEdit = withinTimeLimit && notRead && !m.IsDeleted
	m.CanDelete = withinTimeLimit && notRead && !m.IsDeleted

	return nil
}
