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
	r.POST("/users", handler.POST_DEFAULT_TableDataHandler(database.DB, &model.User{}, []string{"UserRole"}))
	r.PATCH("/users", handler.PATCH_DEFAULT_TableDataHandler(database.DB, &model.User{}, []string{"UserRole"}))
	r.PUT("/users", handler.PUT_DEFAULT_TableDataHandler(database.DB, &model.User{}, []string{"UserRole"}))
	r.DELETE("/users", handler.DELETE_DEFAULT_TableDataHandler(database.DB, &model.User{}))

	// Theme endpoints
	r.GET("/themes", handler.GET_DEFAULT_TableDataHandler(database.DB, &model.Theme{}, []string{}))
	r.POST("/themes", handler.POST_DEFAULT_TableDataHandler(database.DB, &model.Theme{}, []string{}))
	r.PATCH("/themes", handler.PATCH_DEFAULT_TableDataHandler(database.DB, &model.Theme{}, []string{}))
	r.DELETE("/themes", handler.DELETE_DEFAULT_TableDataHandler(database.DB, &model.Theme{}))

	// Product endpoints - Public Read, Protected CUD
	r.GET("/products", handler.GetProducts(database.DB))          // Public: Get all products with filters
	r.GET("/products/:id", handler.GetProductByID(database.DB))   // Public: Get single product
	r.POST("/products", handler.CreateProduct(database.DB))       // Protected: Create product
	r.PUT("/products/:id", handler.UpdateProduct(database.DB))    // Protected: Update product
	r.PATCH("/products/:id", handler.UpdateProduct(database.DB))  // Protected: Update product (alias)
	r.DELETE("/products/:id", handler.DeleteProduct(database.DB)) // Protected: Delete product

	// Category endpoints - Public
	r.GET("/categories", handler.GetCategories(database.DB))        // Get all categories
	r.GET("/sub-categories", handler.GetSubCategories(database.DB)) // Get subcategories

	// Shop endpoints - Public
	r.GET("/shops", handler.GetShops(database.DB)) // Get all shops

	// My Shop endpoints - Protected (User's own shop management)
	r.GET("/my-shop", handler.GetMyShop(database.DB))                   // Get authenticated user's shop
	r.POST("/my-shop", handler.CreateMyShop(database.DB))               // Create shop (upgrade to seller)
	r.PUT("/my-shop", handler.UpdateMyShop(database.DB))                // Update user's shop
	r.PATCH("/my-shop", handler.UpdateMyShop(database.DB))              // Update user's shop (alias)
	r.GET("/my-shop/products", handler.GetMyShopProducts(database.DB))  // Get user's shop products
	r.GET("/my-shop/check", handler.CheckShopAvailability(database.DB)) // Check if user can create shop

	// Wishlist endpoints - Protected (User's wishlist for saved products)
	wishlistHandler := handler.NewWishlistHandler(database.DB)
	r.GET("/my-wishlist", wishlistHandler.GetMyWishlist)          // Get user's wishlist
	r.POST("/wishlist/add", wishlistHandler.AddToWishlist)        // Add product to wishlist
	r.PUT("/wishlist/:id", wishlistHandler.UpdateWishlistItem)    // Update wishlist item notes
	r.PATCH("/wishlist/:id", wishlistHandler.UpdateWishlistItem)  // Update wishlist item (alias)
	r.DELETE("/wishlist/:id", wishlistHandler.RemoveFromWishlist) // Remove item from wishlist
	r.DELETE("/wishlist/clear", wishlistHandler.ClearWishlist)    // Clear entire wishlist

	// Chat endpoints - Protected (User messaging system)
	chatHandler := handler.NewChatHandler(database.DB)
	r.GET("/chats", chatHandler.GetMyChats)                          // Get all user's chats
	r.POST("/chats", chatHandler.GetOrCreateChat)                    // Get or create chat with another user
	r.GET("/chats/:id/messages", chatHandler.GetChatMessages)        // Get messages in a chat
	r.POST("/chats/:id/messages", chatHandler.SendMessage)           // Send message in a chat
	r.PUT("/chats/:id/read", chatHandler.MarkChatRead)               // Mark all messages in chat as read
	r.PUT("/messages/:id/received", chatHandler.MarkMessageReceived) // Mark message as received
	r.PUT("/messages/:id/read", chatHandler.MarkMessageRead)         // Mark message as read
	r.PUT("/messages/:id", chatHandler.EditMessage)                  // Edit message (within 7 mins, not read)
	r.DELETE("/messages/:id", chatHandler.DeleteMessage)             // Delete message (within 7 mins, not read)
	r.GET("/messages/unread", chatHandler.GetUnreadMessages)         // Get all unread messages
	r.GET("/messages/unread/count", chatHandler.GetUnreadCount)      // Get unread count

}
