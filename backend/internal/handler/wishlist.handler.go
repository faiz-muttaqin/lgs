package handler

import (
	"net/http"
	"strconv"

	"github.com/faiz-muttaqin/shadcn-admin-go-starter/backend/internal/helper"
	"github.com/faiz-muttaqin/shadcn-admin-go-starter/backend/internal/model"
	"github.com/faiz-muttaqin/shadcn-admin-go-starter/backend/pkg/audit"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type WishlistHandler struct {
	DB *gorm.DB
}

func NewWishlistHandler(db *gorm.DB) *WishlistHandler {
	return &WishlistHandler{DB: db}
}

// GetMyWishlist retrieves the current user's wishlist items
func (h *WishlistHandler) GetMyWishlist(c *gin.Context) {
	db := h.DB
	userData, err := helper.GetFirebaseUser(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"success": false,
			"message": "Unauthorized",
		})
		return
	}

	var wishlistItems []model.WishlistItem
	if err := db.Where("user_id = ?", userData.ID).
		Preload("Product").
		Preload("Product.Shop").
		Preload("Product.Category").
		Preload("Product.SubCategory").
		Preload("Product.Images").
		Order("created_at DESC").
		Find(&wishlistItems).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "Failed to retrieve wishlist",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data": gin.H{
			"items": wishlistItems,
			"count": len(wishlistItems),
		},
	})
}

// AddToWishlist adds a product to wishlist
func (h *WishlistHandler) AddToWishlist(c *gin.Context) {
	db := h.DB
	userData, err := helper.GetFirebaseUser(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"success": false,
			"message": "Unauthorized",
		})
		return
	}

	var input struct {
		ProductID uint   `json:"product_id" binding:"required"`
		Notes     string `json:"notes"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Invalid input",
		})
		return
	}

	// Check if already in wishlist
	var existing model.WishlistItem
	if err := db.Where("user_id = ? AND product_id = ?", userData.ID, input.ProductID).First(&existing).Error; err == nil {
		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"message": "Product already in wishlist",
			"data":    existing,
		})
		return
	}

	wishlistItem := model.WishlistItem{
		UserID:    userData.ID,
		ProductID: input.ProductID,
		Notes:     input.Notes,
	}

	if err := db.Create(&wishlistItem).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "Failed to add to wishlist",
		})
		return
	}

	db.Preload("Product").First(&wishlistItem, wishlistItem.ID)
	audit.Log(c, db, userData, audit.Create("wishlist_item", wishlistItem.ID).After(wishlistItem).Success("Added to wishlist"))

	c.JSON(http.StatusCreated, gin.H{
		"success": true,
		"message": "Added to wishlist",
		"data":    wishlistItem,
	})
}

// UpdateWishlistItem updates wishlist item notes
func (h *WishlistHandler) UpdateWishlistItem(c *gin.Context) {
	db := h.DB
	userData, err := helper.GetFirebaseUser(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"success": false,
			"message": "Unauthorized",
		})
		return
	}

	itemID, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	var input struct {
		Notes string `json:"notes"`
	}
	c.ShouldBindJSON(&input)

	var item model.WishlistItem
	if err := db.Where("id = ? AND user_id = ?", itemID, userData.ID).First(&item).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"success": false,
			"message": "Item not found",
		})
		return
	}

	old := item
	item.Notes = input.Notes
	db.Save(&item)

	audit.Log(c, db, userData, audit.Update("wishlist_item", item.ID).Before(old).After(item).Success("Updated wishlist item"))

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Updated",
		"data":    item,
	})
}

// RemoveFromWishlist removes item from wishlist
func (h *WishlistHandler) RemoveFromWishlist(c *gin.Context) {
	db := h.DB
	userData, err := helper.GetFirebaseUser(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"success": false,
			"message": "Unauthorized",
		})
		return
	}

	itemID, _ := strconv.ParseUint(c.Param("id"), 10, 32)

	var item model.WishlistItem
	if err := db.Where("id = ? AND user_id = ?", itemID, userData.ID).First(&item).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"success": false,
			"message": "Item not found",
		})
		return
	}

	db.Delete(&item)
	audit.Log(c, db, userData, audit.Delete("wishlist_item", item.ID).Before(item).Success("Removed from wishlist"))

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Removed from wishlist",
	})
}

// ClearWishlist clears all wishlist items
func (h *WishlistHandler) ClearWishlist(c *gin.Context) {
	db := h.DB
	userData, err := helper.GetFirebaseUser(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"success": false,
			"message": "Unauthorized",
		})
		return
	}

	var items []model.WishlistItem
	db.Where("user_id = ?", userData.ID).Find(&items)

	if len(items) == 0 {
		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"message": "Wishlist already empty",
		})
		return
	}

	db.Where("user_id = ?", userData.ID).Delete(&model.WishlistItem{})
	audit.Log(c, db, userData, audit.Delete("wishlist", userData.ID).Before(gin.H{"count": len(items)}).Success("Cleared wishlist"))

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Wishlist cleared",
	})
}
