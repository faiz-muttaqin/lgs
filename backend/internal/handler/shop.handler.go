package handler

import (
	"net/http"
	"strings"

	"github.com/faiz-muttaqin/shadcn-admin-go-starter/backend/internal/helper"
	"github.com/faiz-muttaqin/shadcn-admin-go-starter/backend/internal/model"
	"github.com/faiz-muttaqin/shadcn-admin-go-starter/backend/pkg/audit"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// GetMyShop - Get authenticated user's shop
func GetMyShop(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Authentication check
		userData, err := helper.GetFirebaseUser(c)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"success": false,
				"message": "Unauthorized: user authentication failed",
				"error":   err.Error(),
			})
			return
		}

		var shop model.Shop
		if err := db.Where("user_id = ?", userData.ID).First(&shop).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				c.JSON(http.StatusNotFound, gin.H{
					"success":  false,
					"message":  "You don't have a shop yet. Please create one to start selling.",
					"has_shop": false,
				})
				return
			}
			c.JSON(http.StatusInternalServerError, gin.H{
				"success": false,
				"message": "Failed to fetch shop",
				"error":   err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"success":  true,
			"message":  "Shop fetched successfully",
			"data":     shop,
			"has_shop": true,
		})
	}
}

// CreateMyShop - Create shop for authenticated user (upgrade to seller)
func CreateMyShop(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Authentication check
		userData, err := helper.GetFirebaseUser(c)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"success": false,
				"message": "Unauthorized: user authentication failed",
				"error":   err.Error(),
			})
			return
		}

		// Check if user already has a shop
		var existingShop model.Shop
		if err := db.Where("user_id = ?", userData.ID).First(&existingShop).Error; err == nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"success": false,
				"message": "You already have a shop",
				"data":    existingShop,
			})
			return
		}

		var shopData struct {
			Name        string `json:"name" binding:"required"`
			City        string `json:"city" binding:"required"`
			Description string `json:"description"`
			ImageURL    string `json:"image_url"`
			Domain      string `json:"domain"`
		}

		if err := c.ShouldBindJSON(&shopData); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"success": false,
				"message": "Invalid request body. Required fields: name, city",
				"error":   err.Error(),
			})
			return
		}

		// Generate slug from shop name
		slug := generateShopSlug(shopData.Name)

		// Create shop
		shop := model.Shop{
			UserID:      userData.ID,
			Name:        shopData.Name,
			Slug:        slug,
			City:        shopData.City,
			Description: shopData.Description,
			ImageURL:    shopData.ImageURL,
			Domain:      shopData.Domain,
			IsOfficial:  false, // Default to false, admin can upgrade
			IsActive:    true,
		}

		if err := db.Create(&shop).Error; err != nil {
			// Log failed creation
			audit.Log(c, db, userData,
				audit.Create("shop", shop.ID).After(shop).Failed(err),
			)
			c.JSON(http.StatusInternalServerError, gin.H{
				"success": false,
				"message": "Failed to create shop",
				"error":   err.Error(),
			})
			return
		}

		// Update user role to seller if not already
		if userData.RoleID == 2 { // Default/Customer role
			db.Model(&userData).Update("role_id", 3) // Seller role
		}

		// Log successful creation
		audit.Log(c, db, userData,
			audit.Create("shop", shop.ID).After(shop).Success("Shop created successfully and user upgraded to seller"),
		)

		c.JSON(http.StatusCreated, gin.H{
			"success": true,
			"message": "Shop created successfully! You are now a seller.",
			"data":    shop,
		})
	}
}

// UpdateMyShop - Update authenticated user's shop
func UpdateMyShop(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Authentication check
		userData, err := helper.GetFirebaseUser(c)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"success": false,
				"message": "Unauthorized: user authentication failed",
				"error":   err.Error(),
			})
			return
		}

		// Get user's shop
		var shop model.Shop
		if err := db.Where("user_id = ?", userData.ID).First(&shop).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				c.JSON(http.StatusNotFound, gin.H{
					"success": false,
					"message": "You don't have a shop. Please create one first.",
				})
				return
			}
			c.JSON(http.StatusInternalServerError, gin.H{
				"success": false,
				"message": "Failed to fetch shop",
				"error":   err.Error(),
			})
			return
		}

		var updateData struct {
			Name        string `json:"name"`
			City        string `json:"city"`
			Description string `json:"description"`
			ImageURL    string `json:"image_url"`
			Domain      string `json:"domain"`
		}

		if err := c.ShouldBindJSON(&updateData); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"success": false,
				"message": "Invalid request body",
				"error":   err.Error(),
			})
			return
		}

		// Store old shop data for audit
		oldShop := shop

		// Update shop
		if err := db.Model(&shop).Updates(updateData).Error; err != nil {
			// Log failed update
			audit.Log(c, db, userData,
				audit.Update("shop", shop.ID).Before(oldShop).After(shop).Failed(err),
			)
			c.JSON(http.StatusInternalServerError, gin.H{
				"success": false,
				"message": "Failed to update shop",
				"error":   err.Error(),
			})
			return
		}

		// Reload shop
		db.First(&shop, shop.ID)

		// Log successful update
		audit.Log(
			c,
			db,
			userData,
			audit.Update("shop", shop.ID).
				Before(oldShop).
				After(shop).
				Success("Shop updated successfully"),
		)

		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"message": "Shop updated successfully",
			"data":    shop,
		})
	}
}

// GetMyShopProducts - Get products from authenticated user's shop
func GetMyShopProducts(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Authentication check
		userData, err := helper.GetFirebaseUser(c)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"success": false,
				"message": "Unauthorized: user authentication failed",
				"error":   err.Error(),
			})
			return
		}

		// Get user's shop
		var shop model.Shop
		if err := db.Where("user_id = ?", userData.ID).First(&shop).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{
				"success": false,
				"message": "You don't have a shop yet",
			})
			return
		}

		// Get products
		var products []model.Product
		query := db.Where("shop_id = ?", shop.ID).
			Preload("Category").Preload("SubCategory").
			Preload("Images").Preload("Labels").Preload("Badges").Preload("Variants")

		if err := query.Find(&products).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"success": false,
				"message": "Failed to fetch products",
				"error":   err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"message": "Products fetched successfully",
			"data":    products,
			"shop":    shop,
		})
	}
}

// CheckShopAvailability - Check if user can create a shop
func CheckShopAvailability(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Authentication check
		userData, err := helper.GetFirebaseUser(c)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"success": false,
				"message": "Unauthorized: user authentication failed",
				"error":   err.Error(),
			})
			return
		}

		// Check if user has a shop
		var shop model.Shop
		hasShop := db.Where("user_id = ?", userData.ID).First(&shop).Error == nil

		c.JSON(http.StatusOK, gin.H{
			"success":    true,
			"has_shop":   hasShop,
			"can_create": !hasShop,
			"user_role":  userData.UserRole.Name,
			"message": func() string {
				if hasShop {
					return "You already have a shop"
				}
				return "You can create a shop"
			}(),
		})
	}
}

// Helper function to generate shop slug
func generateShopSlug(name string) string {
	slug := strings.ToLower(name)
	slug = strings.ReplaceAll(slug, " ", "-")
	slug = strings.ReplaceAll(slug, "_", "-")
	// Remove special characters (basic implementation)
	// In production, use a proper slug library
	return slug
}
