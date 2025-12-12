package handler

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/faiz-muttaqin/shadcn-admin-go-starter/backend/internal/helper"
	"github.com/faiz-muttaqin/shadcn-admin-go-starter/backend/internal/model"
	"github.com/faiz-muttaqin/shadcn-admin-go-starter/backend/pkg/audit"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// GetProducts - Public endpoint to read products with pagination, filtering, sorting
func GetProducts(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var products []model.Product
		query := db.Model(&model.Product{})

		// Preload relations
		query = query.Preload("Category").Preload("SubCategory").Preload("Shop").
			Preload("Images").Preload("Labels").Preload("Badges").Preload("Variants")

		// Filtering
		if categoryID := c.Query("category_id"); categoryID != "" {
			query = query.Where("category_id = ?", categoryID)
		}
		if subCategoryID := c.Query("sub_category_id"); subCategoryID != "" {
			query = query.Where("sub_category_id = ?", subCategoryID)
		}
		if shopID := c.Query("shop_id"); shopID != "" {
			query = query.Where("shop_id = ?", shopID)
		}
		if isActive := c.Query("is_active"); isActive != "" {
			query = query.Where("is_active = ?", isActive)
		}
		if isFeatured := c.Query("is_featured"); isFeatured != "" {
			query = query.Where("is_featured = ?", isFeatured)
		}
		if status := c.Query("status"); status != "" {
			query = query.Where("status = ?", status)
		}
		if search := c.Query("search"); search != "" {
			query = query.Where("name LIKE ? OR description LIKE ?", "%"+search+"%", "%"+search+"%")
		}

		// Price range filtering
		if minPrice := c.Query("min_price"); minPrice != "" {
			query = query.Where("price >= ?", minPrice)
		}
		if maxPrice := c.Query("max_price"); maxPrice != "" {
			query = query.Where("price <= ?", maxPrice)
		}

		// Sorting
		sortBy := c.DefaultQuery("sort_by", "created_at")
		sortOrder := c.DefaultQuery("sort_order", "desc")
		query = query.Order(fmt.Sprintf("%s %s", sortBy, sortOrder))

		// Pagination
		page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
		pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "20"))
		offset := (page - 1) * pageSize

		// Count total
		var total int64
		query.Count(&total)

		// Execute query with pagination
		if err := query.Offset(offset).Limit(pageSize).Find(&products).Error; err != nil {
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
			"meta": gin.H{
				"page":       page,
				"page_size":  pageSize,
				"total":      total,
				"total_page": (total + int64(pageSize) - 1) / int64(pageSize),
			},
		})
	}
}

// GetProductByID - Public endpoint to get single product by ID
func GetProductByID(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		var product model.Product

		if err := db.Preload("Category").Preload("SubCategory").Preload("Shop").
			Preload("Images").Preload("Labels").Preload("Badges").Preload("Variants").
			First(&product, id).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				c.JSON(http.StatusNotFound, gin.H{
					"success": false,
					"message": "Product not found",
				})
				return
			}
			c.JSON(http.StatusInternalServerError, gin.H{
				"success": false,
				"message": "Failed to fetch product",
				"error":   err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"message": "Product fetched successfully",
			"data":    product,
		})
	}
}

// CreateProduct - Protected endpoint to create a new product
func CreateProduct(db *gorm.DB) gin.HandlerFunc {
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
		fmt.Println("User creating product:", userData.Email)

		// Check if user has a shop
		var userShop model.Shop
		if err := db.Where("user_id = ?", userData.ID).First(&userShop).Error; err != nil {
			c.JSON(http.StatusForbidden, gin.H{
				"success":  false,
				"message":  "You must create a shop first before adding products. Please visit /my-shop to create your shop.",
				"has_shop": false,
			})
			return
		}

		var product model.Product
		if err := c.ShouldBindJSON(&product); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"success": false,
				"message": "Invalid request body",
				"error":   err.Error(),
			})
			return
		}

		// Automatically set the shop_id to user's shop
		product.ShopID = userShop.ID

		// Generate slug from name if not provided
		if product.Slug == "" {
			product.Slug = generateSlug(product.Name)
		}

		// Set default status if not provided
		if product.Status == "" {
			product.Status = model.ProductStatusDraft
		}

		if err := db.Create(&product).Error; err != nil {
			// Log failed creation
			audit.Log(
				c,
				db,
				userData,
				audit.Create("product", product.ID).
					After(product).
					Failed(err),
			)
			c.JSON(http.StatusInternalServerError, gin.H{
				"success": false,
				"message": "Failed to create product",
				"error":   err.Error(),
			})
			return
		}

		// Log successful creation
		audit.Log(
			c,
			db,
			userData,
			audit.Create("product", product.ID).
				After(product).
				Success("Product created successfully"),
		)

		c.JSON(http.StatusCreated, gin.H{
			"success": true,
			"message": "Product created successfully",
			"data":    product,
		})
	}
}

// UpdateProduct - Protected endpoint to update a product
func UpdateProduct(db *gorm.DB) gin.HandlerFunc {
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
		fmt.Println("User updating product:", userData.Email)

		id := c.Param("id")
		var product model.Product

		// Check if product exists
		if err := db.Preload("Shop").First(&product, id).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				c.JSON(http.StatusNotFound, gin.H{
					"success": false,
					"message": "Product not found",
				})
				return
			}
			c.JSON(http.StatusInternalServerError, gin.H{
				"success": false,
				"message": "Failed to fetch product",
				"error":   err.Error(),
			})
			return
		}

		// Check if product belongs to user's shop (unless super admin)
		if userData.RoleID != 1 { // Not super admin
			var userShop model.Shop
			if err := db.Where("user_id = ?", userData.ID).First(&userShop).Error; err != nil {
				c.JSON(http.StatusForbidden, gin.H{
					"success": false,
					"message": "You don't have a shop",
				})
				return
			}
			if product.ShopID != userShop.ID {
				c.JSON(http.StatusForbidden, gin.H{
					"success": false,
					"message": "You can only update products from your own shop",
				})
				return
			}
		}

		var updateData model.Product
		if err := c.ShouldBindJSON(&updateData); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"success": false,
				"message": "Invalid request body",
				"error":   err.Error(),
			})
			return
		}

		// Store old product data for audit
		oldProduct := product

		// Update product
		if err := db.Model(&product).Updates(updateData).Error; err != nil {
			// Log failed update
			audit.Log(
				c,
				db,
				userData,
				audit.Update("product", product.ID).
					Before(oldProduct).
					After(product).
					Failed(err),
			)
			c.JSON(http.StatusInternalServerError, gin.H{
				"success": false,
				"message": "Failed to update product",
				"error":   err.Error(),
			})
			return
		}

		// Reload product with relations
		db.Preload("Category").Preload("SubCategory").Preload("Shop").
			Preload("Images").Preload("Labels").Preload("Badges").Preload("Variants").
			First(&product, id)

		// Log successful update
		audit.Log(
			c,
			db,
			userData,
			audit.Update("product", product.ID).
				Before(oldProduct).
				After(product).
				Success("Product updated successfully"),
		)

		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"message": "Product updated successfully",
			"data":    product,
		})
	}
}

// DeleteProduct - Protected endpoint to delete a product (soft delete)
func DeleteProduct(db *gorm.DB) gin.HandlerFunc {
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
		fmt.Println("User deleting product:", userData.Email)

		id := c.Param("id")
		var product model.Product

		// Check if product exists
		if err := db.Preload("Shop").First(&product, id).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				c.JSON(http.StatusNotFound, gin.H{
					"success": false,
					"message": "Product not found",
				})
				return
			}
			c.JSON(http.StatusInternalServerError, gin.H{
				"success": false,
				"message": "Failed to fetch product",
				"error":   err.Error(),
			})
			return
		}

		// Check if product belongs to user's shop (unless super admin)
		if userData.RoleID != 1 { // Not super admin
			var userShop model.Shop
			if err := db.Where("user_id = ?", userData.ID).First(&userShop).Error; err != nil {
				c.JSON(http.StatusForbidden, gin.H{
					"success": false,
					"message": "You don't have a shop",
				})
				return
			}
			if product.ShopID != userShop.ID {
				c.JSON(http.StatusForbidden, gin.H{
					"success": false,
					"message": "You can only delete products from your own shop",
				})
				return
			}
		}

		// Store product data for audit before deletion
		productBeforeDelete := product

		// Soft delete
		if err := db.Delete(&product).Error; err != nil {
			// Log failed deletion
			audit.Log(
				c,
				db,
				userData,
				audit.Delete("product", product.ID).
					Before(productBeforeDelete).
					Failed(err),
			)
			c.JSON(http.StatusInternalServerError, gin.H{
				"success": false,
				"message": "Failed to delete product",
				"error":   err.Error(),
			})
			return
		}

		// Log successful deletion
		audit.Log(
			c,
			db,
			userData,
			audit.Delete("product", product.ID).
				Before(productBeforeDelete).
				Success("Product deleted successfully"),
		)

		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"message": "Product deleted successfully",
		})
	}
}

// Helper functions

// GetCategories - Public endpoint to get all categories
func GetCategories(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var categories []model.Category
		if err := db.Preload("SubCategories").Where("is_active = ?", true).Find(&categories).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"success": false,
				"message": "Failed to fetch categories",
				"error":   err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"message": "Categories fetched successfully",
			"data":    categories,
		})
	}
}

// GetSubCategories - Public endpoint to get subcategories
func GetSubCategories(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var subCategories []model.SubCategory
		query := db.Model(&model.SubCategory{}).Preload("Category")

		if categoryID := c.Query("category_id"); categoryID != "" {
			query = query.Where("category_id = ?", categoryID)
		}

		query = query.Where("is_active = ?", true)

		if err := query.Find(&subCategories).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"success": false,
				"message": "Failed to fetch subcategories",
				"error":   err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"message": "Subcategories fetched successfully",
			"data":    subCategories,
		})
	}
}

// GetShops - Public endpoint to get all shops
func GetShops(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var shops []model.Shop
		if err := db.Find(&shops).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"success": false,
				"message": "Failed to fetch shops",
				"error":   err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"message": "Shops fetched successfully",
			"data":    shops,
		})
	}
}

// generateSlug creates a URL-friendly slug from a string
func generateSlug(name string) string {
	slug := strings.ToLower(name)
	slug = strings.ReplaceAll(slug, " ", "-")
	// Remove special characters (basic implementation)
	return slug
}
