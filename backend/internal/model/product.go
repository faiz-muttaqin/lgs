package model

import (
	"time"

	"github.com/faiz-muttaqin/lgs/backend/pkg/types"
	"gorm.io/gorm"
)

// Category represents the main product category (e.g., "Dapur", "Elektronik")
type Category struct {
	ID          uint           `gorm:"primaryKey;column:id" json:"id"`
	Name        string         `gorm:"column:name;size:100;unique;not null" json:"name" ui:"creatable;visible;editable;filterable;sortable"`
	Slug        string         `gorm:"column:slug;size:100;unique;not null;index" json:"slug" ui:"visible;filterable"`
	Description string         `gorm:"column:description;type:text" json:"description" ui:"creatable;visible;editable"`
	ImageURL    string         `gorm:"column:image_url;size:500" json:"image_url" ui:"creatable;visible;editable"`
	IsActive    bool           `gorm:"column:is_active;default:true" json:"is_active" ui:"visible;editable;filterable"`
	CreatedAt   time.Time      `gorm:"column:created_at" json:"created_at"`
	UpdatedAt   time.Time      `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`

	// Relations
	SubCategories []SubCategory `gorm:"foreignKey:CategoryID" json:"sub_categories,omitempty"`
}

func (Category) TableName() string {
	return "categories"
}

// SubCategory represents product subcategory (e.g., "Alat Masak Khusus")
type SubCategory struct {
	ID          uint           `gorm:"primaryKey;column:id" json:"id"`
	CategoryID  uint           `gorm:"column:category_id;not null;index" json:"category_id" ui:"creatable;editable;filterable;selection:/options?data=category"`
	Name        string         `gorm:"column:name;size:100;not null" json:"name" ui:"creatable;visible;editable;filterable;sortable"`
	Slug        string         `gorm:"column:slug;size:100;not null;index" json:"slug" ui:"visible;filterable"`
	Description string         `gorm:"column:description;type:text" json:"description" ui:"creatable;visible;editable"`
	ImageURL    string         `gorm:"column:image_url;size:500" json:"image_url" ui:"creatable;visible;editable"`
	IsActive    bool           `gorm:"column:is_active;default:true" json:"is_active" ui:"visible;editable;filterable"`
	CreatedAt   time.Time      `gorm:"column:created_at" json:"created_at"`
	UpdatedAt   time.Time      `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`

	// Relations
	Category Category  `gorm:"foreignKey:CategoryID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"category,omitempty"`
	Products []Product `gorm:"foreignKey:SubCategoryID" json:"products,omitempty"`
}

func (SubCategory) TableName() string {
	return "sub_categories"
}

// Shop represents the seller/shop information
type Shop struct {
	// ExternalID  string         `gorm:"column:external_id;size:100;unique" json:"external_id"`
	ID          uint           `gorm:"primaryKey;column:id" json:"id"`
	UserID      uint           `gorm:"column:user_id;unique;not null;index" json:"user_id" ui:"visible;filterable"`
	Name        string         `gorm:"column:name;size:200;not null" json:"name" ui:"creatable;visible;editable;filterable;sortable"`
	Slug        string         `gorm:"column:slug;size:200;unique;not null;index" json:"slug"`
	Domain      string         `gorm:"column:domain;size:200" json:"domain"`
	City        string         `gorm:"column:city;size:100" json:"city" ui:"visible;filterable"`
	ImageURL    string         `gorm:"column:image_url;size:500" json:"image_url"`
	Reputation  string         `gorm:"column:reputation;size:500" json:"reputation"`
	IsOfficial  bool           `gorm:"column:is_official;default:false" json:"is_official" ui:"visible;filterable"`
	IsActive    bool           `gorm:"column:is_active;default:true" json:"is_active" ui:"visible;filterable"`
	Description string         `gorm:"column:description;type:text" json:"description" ui:"creatable;visible;editable"`
	CreatedAt   time.Time      `gorm:"column:created_at" json:"created_at"`
	UpdatedAt   time.Time      `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`

	// Relations
	Products []Product `gorm:"foreignKey:ShopID" json:"products,omitempty"`
}

func (Shop) TableName() string {
	return "shops"
}

// BeforeCreate hook to generate slug if not provided
func (s *Shop) BeforeCreate(tx *gorm.DB) error {
	if s.Slug == "" {
		s.Slug = generateShopSlug(s.Name)
	}
	return nil
}

// Helper function to generate shop slug
func generateShopSlug(name string) string {
	slug := name
	// Simple implementation - in production use a proper slug library
	return slug
}

// Product represents the main product model
type Product struct {
	ID            uint           `gorm:"primaryKey;column:id" json:"id" ui:"sortable,filterable"`
	ExternalID    string         `gorm:"column:external_id;size:100" json:"external_id"`
	SKU           string         `gorm:"column:sku;size:100;unique;index" json:"sku" ui:"creatable;visible;editable;filterable"`
	Name          string         `gorm:"column:name;size:255;not null" json:"name" ui:"creatable;visible;editable;filterable;sortable"`
	Slug          string         `gorm:"column:slug;size:255;unique;not null;index" json:"slug"`
	Subtitle      string         `gorm:"column:subtitle;size:255" json:"subtitle" ui:"creatable;visible;editable"`
	Description   string         `gorm:"column:description;type:text" json:"description" ui:"creatable;visible;editable"`
	ImageURL      string         `gorm:"column:image_url;size:500;not null" json:"image_url" ui:"creatable;visible;editable"`
	Price         float64        `gorm:"column:price;not null" json:"price" ui:"creatable;visible;editable;filterable;sortable"`
	SlashedPrice  float64        `gorm:"column:slashed_price" json:"slashed_price" ui:"creatable;visible;editable"`
	DiscountPct   int            `gorm:"column:discount_pct;default:0" json:"discount_pct" ui:"visible;filterable"`
	Stock         int            `gorm:"column:stock;default:0" json:"stock" ui:"creatable;visible;editable;filterable;sortable"`
	Rating        float32        `gorm:"column:rating;default:0" json:"rating" ui:"visible;filterable;sortable"`
	CountReview   int            `gorm:"column:count_review;default:0" json:"count_review" ui:"visible;sortable"`
	CountSold     int            `gorm:"column:count_sold;default:0" json:"count_sold" ui:"visible;sortable"`
	Weight        int            `gorm:"column:weight;comment:in grams" json:"weight" ui:"creatable;visible;editable"`
	IsActive      bool           `gorm:"column:is_active;default:true" json:"is_active" ui:"visible;editable;filterable"`
	IsFeatured    bool           `gorm:"column:is_featured;default:false" json:"is_featured" ui:"visible;editable;filterable"`
	Status        types.Badge    `gorm:"column:status;default:'draft'" json:"status" ui:"visible;editable;filterable;sortable;selection:/options?data=product_status"`
	CategoryID    uint           `gorm:"column:category_id;index" json:"category_id" ui:"creatable;editable;filterable;selection:/options?data=category"`
	SubCategoryID uint           `gorm:"column:sub_category_id;index" json:"sub_category_id" ui:"creatable;editable;filterable;selection:/options?data=sub_category"`
	ShopID        uint           `gorm:"column:shop_id;not null;index" json:"shop_id" ui:"creatable;editable;filterable;selection:/options?data=shop"`
	CreatedAt     time.Time      `gorm:"column:created_at" json:"created_at"`
	UpdatedAt     time.Time      `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt     gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`

	// Relations
	Category    Category         `gorm:"foreignKey:CategoryID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL" json:"category,omitempty"`
	SubCategory SubCategory      `gorm:"foreignKey:SubCategoryID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL" json:"sub_category,omitempty"`
	Shop        Shop             `gorm:"foreignKey:ShopID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"shop,omitempty"`
	Images      []ProductImage   `gorm:"foreignKey:ProductID" json:"images,omitempty"`
	Labels      []ProductLabel   `gorm:"foreignKey:ProductID" json:"labels,omitempty"`
	Badges      []ProductBadge   `gorm:"foreignKey:ProductID" json:"badges,omitempty"`
	Variants    []ProductVariant `gorm:"foreignKey:ProductID" json:"variants,omitempty"`
}

func (Product) TableName() string {
	return "products"
}

// ProductImage represents additional product images
type ProductImage struct {
	ID        uint           `gorm:"primaryKey;column:id" json:"id"`
	ProductID uint           `gorm:"column:product_id;not null;index" json:"product_id"`
	ImageURL  string         `gorm:"column:image_url;size:500;not null" json:"image_url"`
	Position  int            `gorm:"column:position;default:0" json:"position"`
	IsMain    bool           `gorm:"column:is_main;default:false" json:"is_main"`
	CreatedAt time.Time      `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time      `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`

	// Relations
	Product Product `gorm:"foreignKey:ProductID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"-"`
}

func (ProductImage) TableName() string {
	return "product_images"
}

// ProductLabel represents labels like "Flash Sale", "Promo Guncang", etc.
type ProductLabel struct {
	ID        uint           `gorm:"primaryKey;column:id" json:"id"`
	ProductID uint           `gorm:"column:product_id;not null;index" json:"product_id"`
	Title     string         `gorm:"column:title;size:100;not null" json:"title"`
	Type      string         `gorm:"column:type;size:50" json:"type"`
	Position  string         `gorm:"column:position;size:50" json:"position"`
	URL       string         `gorm:"column:url;size:500" json:"url"`
	CreatedAt time.Time      `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time      `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`

	// Relations
	Product Product `gorm:"foreignKey:ProductID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"-"`
}

func (ProductLabel) TableName() string {
	return "product_labels"
}

// ProductBadge represents shop badges like "Official Store"
type ProductBadge struct {
	ID        uint           `gorm:"primaryKey;column:id" json:"id"`
	ProductID uint           `gorm:"column:product_id;not null;index" json:"product_id"`
	Title     string         `gorm:"column:title;size:100;not null" json:"title"`
	ImageURL  string         `gorm:"column:image_url;size:500" json:"image_url"`
	CreatedAt time.Time      `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time      `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`

	// Relations
	Product Product `gorm:"foreignKey:ProductID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"-"`
}

func (ProductBadge) TableName() string {
	return "product_badges"
}

// ProductVariant represents product variations (size, color, etc.)
type ProductVariant struct {
	ID          uint           `gorm:"primaryKey;column:id" json:"id"`
	ProductID   uint           `gorm:"column:product_id;not null;index" json:"product_id"`
	SKU         string         `gorm:"column:sku;size:100;unique;index" json:"sku"`
	Name        string         `gorm:"column:name;size:100;not null" json:"name"`
	Price       float64        `gorm:"column:price" json:"price"`
	Stock       int            `gorm:"column:stock;default:0" json:"stock"`
	Weight      int            `gorm:"column:weight;comment:in grams" json:"weight"`
	ImageURL    string         `gorm:"column:image_url;size:500" json:"image_url"`
	IsAvailable bool           `gorm:"column:is_available;default:true" json:"is_available"`
	CreatedAt   time.Time      `gorm:"column:created_at" json:"created_at"`
	UpdatedAt   time.Time      `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`

	// Relations
	Product Product `gorm:"foreignKey:ProductID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"-"`
}

func (ProductVariant) TableName() string {
	return "product_variants"
}

// Product status constants
const (
	ProductStatusDraft      = "draft"
	ProductStatusPublished  = "published"
	ProductStatusArchived   = "archived"
	ProductStatusOutOfStock = "out_of_stock"
)
