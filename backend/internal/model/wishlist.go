package model

import (
	"time"

	"gorm.io/gorm"
)

// WishlistItem represents an item in the user's wishlist
// Users can save products they're interested in for later
type WishlistItem struct {
	ID        uint           `gorm:"primaryKey;column:id" json:"id"`
	UserID    uint           `gorm:"column:user_id;not null;uniqueIndex:idx_user_product" json:"user_id"`
	ProductID uint           `gorm:"column:product_id;not null;uniqueIndex:idx_user_product" json:"product_id"`
	Notes     string         `gorm:"column:notes;type:text" json:"notes,omitempty"`
	CreatedAt time.Time      `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time      `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`

	// Relations
	User    User    `gorm:"foreignKey:UserID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"user,omitempty"`
	Product Product `gorm:"foreignKey:ProductID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"product,omitempty"`
}

func (WishlistItem) TableName() string {
	return "wishlist_items"
}
