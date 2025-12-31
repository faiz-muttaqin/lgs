package routes

import (
	"github.com/faiz-muttaqin/lgs/backend/internal/database"
	"github.com/faiz-muttaqin/lgs/backend/internal/handler"
	"github.com/faiz-muttaqin/lgs/backend/internal/model"
	"github.com/faiz-muttaqin/lgs/backend/pkg/util"
	"github.com/gin-gonic/gin"
)

var R *gin.Engine

func Routes() {
	// Endpoint login API
	r := R.Group(util.GetPathOnly(util.Getenv("VITE_BACKEND", "/api")))

	r.GET("/options", handler.GetOptions())
	// r.GET("/users/hehe", handler.GET_DEFAULT_TableDataHandler(database.DB, &model.User{}, []string{"Role"}))
	r.OPTIONS("/auth/login", handler.GetAuthLogin())
	r.GET("/auth/login", handler.GetAuthLogin())
	r.GET("/auth/logout", handler.GetAuthLogout())
	r.GET("/auth/verify", handler.VerifyAuth()) // Test auth endpoint
	r.GET("/roles", handler.GetRoles())
	r.GET("/users", handler.GET_DEFAULT_TableDataHandler(database.DB, &model.User{}, []string{"UserRole"}))
	// r.POST("/users", handler.POST_DEFAULT_TableDataHandler(database.DB, &model.User{}, []string{"UserRole"}))
	// r.PATCH("/users", handler.PATCH_DEFAULT_TableDataHandler(database.DB, &model.User{}, []string{"UserRole"}))
	// r.PUT("/users", handler.PUT_DEFAULT_TableDataHandler(database.DB, &model.User{}, []string{"UserRole"}))
	// r.DELETE("/users", handler.DELETE_DEFAULT_TableDataHandler(database.DB, &model.User{}))

	// Theme endpoints
	r.GET("/themes", handler.GET_DEFAULT_TableDataHandler(database.DB, &model.Theme{}, []string{}))
	r.POST("/themes", handler.POST_DEFAULT_TableDataHandler(database.DB, &model.Theme{}, []string{}))
	r.PATCH("/themes", handler.PATCH_DEFAULT_TableDataHandler(database.DB, &model.Theme{}, []string{}))
	r.DELETE("/themes", handler.DELETE_DEFAULT_TableDataHandler(database.DB, &model.Theme{}))

	// Product endpoints - Public Read, Protected CUD
	r.GET("/products", handler.GET_DEFAULT_TABLE(database.DB, &model.Product{}, []string{"Shop", "Category", "SubCategory", "Images", "Labels", "Badges"})) // Public: Get all products with filters
	r.GET("/products/:id", handler.GetProductByID(database.DB))                                                                                             // Public: Get single product

	r.POST("/products", handler.CreateProduct(database.DB))               // Protected: Create product
	r.GET("/products/batch/template", handler.CreateProduct(database.DB)) // Protected: Create product
	r.POST("/products/batch/upload", handler.CreateProduct(database.DB))  // Protected: Create product
	r.POST("/products/batch/confirm", handler.CreateProduct(database.DB)) // Protected: Create product

	r.PUT("/products/:id", handler.UpdateProduct(database.DB))    // Protected: Update product
	r.PATCH("/products/:id", handler.UpdateProduct(database.DB))  // Protected: Update product (alias)
	r.DELETE("/products/:id", handler.DeleteProduct(database.DB)) // Protected: Delete product

	// Category endpoints - Public
	r.GET("/categories", handler.GET_DEFAULT_TABLE(database.DB, &model.Category{}, []string{}))                               // Get all categories
	r.GET("/categories/sub-categories", handler.GET_DEFAULT_TABLE(database.DB, &model.Category{}, []string{"SubCategories"})) // Get all categories
	r.GET("/sub-categories", handler.GET_DEFAULT_TABLE(database.DB, &model.SubCategory{}, []string{}))                        // Get subcategories
	// r.GET("/categories", handler.GetCategories(database.DB))        // Get all categories
	// r.GET("/sub-categories", handler.GetSubCategories(database.DB)) // Get subcategories

	// Wishlist endpoints - Protected (User's wishlist for saved products)
	r.GET("/my-wishlist", handler.GetMyWishlist(database.DB))          // Get user's wishlist
	r.POST("/wishlist/add", handler.AddToWishlist(database.DB))        // Add product to wishlist
	r.PUT("/wishlist/:id", handler.UpdateWishlistItem(database.DB))    // Update wishlist item notes
	r.PATCH("/wishlist/:id", handler.UpdateWishlistItem(database.DB))  // Update wishlist item (alias)
	r.DELETE("/wishlist/:id", handler.RemoveFromWishlist(database.DB)) // Remove item from wishlist
	r.DELETE("/wishlist/clear", handler.ClearWishlist(database.DB))    // Clear entire wishlist

	// Chat endpoints - Protected (User messaging system)
	r.GET("/chats", handler.GetMyChats(database.DB))                          // Get all user's chats
	r.POST("/chats", handler.GetOrCreateChat(database.DB))                    // Get or create chat with another user
	r.GET("/chats/:id/messages", handler.GetChatMessages(database.DB))        // Get messages in a chat
	r.POST("/chats/:id/messages", handler.SendMessage(database.DB))           // Send message in a chat
	r.PUT("/chats/:id/read", handler.MarkChatRead(database.DB))               // Mark all messages in chat as read
	r.PUT("/messages/:id/received", handler.MarkMessageReceived(database.DB)) // Mark message as received
	r.PUT("/messages/:id/read", handler.MarkMessageRead(database.DB))         // Mark message as read
	r.PUT("/messages/:id", handler.EditMessage(database.DB))                  // Edit message (within 7 mins, not read)
	r.DELETE("/messages/:id", handler.DeleteMessage(database.DB))             // Delete message (within 7 mins, not read)
	r.GET("/messages/unread", handler.GetUnreadMessages(database.DB))         // Get all unread messages
	r.GET("/messages/unread/count", handler.GetUnreadCount(database.DB))      // Get unread count

	// Shop endpoints - Public
	r.GET("/shops", handler.GET_DEFAULT_TABLE(database.DB, &model.Shop{}, []string{}))                    // Get all shops
	r.GET("/shops/products", handler.GET_DEFAULT_TABLE(database.DB, &model.Shop{}, []string{"Products"})) // Get all shops
	// r.GET("/shops", handler.GetShops(database.DB)) // Get all shops

	// My Shop endpoints - Protected (User's own shop management)
	r.GET("/my-shop", handler.GetMyShop(database.DB))                   // Get authenticated user's shop
	r.POST("/my-shop", handler.CreateMyShop(database.DB))               // Create shop (upgrade to seller)
	r.PUT("/my-shop", handler.UpdateMyShop(database.DB))                // Update user's shop
	r.PATCH("/my-shop", handler.UpdateMyShop(database.DB))              // Update user's shop (alias)
	r.GET("/my-shop/products", handler.GetMyShopProducts(database.DB))  // Get user's shop products
	r.GET("/my-shop/check", handler.CheckShopAvailability(database.DB)) // Check if user can create shop

}
