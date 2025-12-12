package database

import (
	"fmt"

	"github.com/faiz-muttaqin/shadcn-admin-go-starter/backend/internal/model"

	"gorm.io/gorm"
)

func AutoMigrateDB(db *gorm.DB) error {
	if err := db.AutoMigrate(
		&model.UserRole{},
		&model.UserAbilityRule{},
		&model.User{},
		&model.WishlistItem{},
	); err != nil {
		return err
	}

	// Pastikan role default tersedia
	db.FirstOrCreate(&model.UserRole{
		ID:    1,
		Title: "Super Admin",
		Name:  "superadmin",
		Icon:  "bx bx-sparkle",
	})

	db.FirstOrCreate(&model.UserRole{
		ID:    2,
		Title: "Default", // costumer role
		Name:  "default",
		Icon:  "bx bx-radio-circle",
	})

	db.FirstOrCreate(&model.UserRole{
		ID:    3,
		Title: "Seller", // seller role
		Name:  "seller",
		Icon:  "bx bx-radio-circle",
	})
	db.FirstOrCreate(&model.UserRole{
		ID:    4,
		Title: "Seller Manager", // management role
		Name:  "seller_manager",
		Icon:  "bx bx-radio-circle",
	})

	// Isi ability rule untuk role default
	var count int64
	db.Model(&model.UserAbilityRule{}).Where("role_id IN ?", []int{1, 2}).Count(&count)
	if count == 0 {
		rules := []model.UserAbilityRule{
			{RoleID: 1, Subject: "*", Read: true},
			{RoleID: 2, Subject: "/", Read: true},
			{RoleID: 2, Subject: "/profile", Read: true, Update: true},
		}
		if err := db.Create(&rules).Error; err != nil {
			return fmt.Errorf("failed creating default abilities: %w", err)
		}
	}

	// Create default users
	var userCount int64
	db.Model(&model.User{}).Count(&userCount)
	if userCount == 0 {
		// Super Admin
		superAdmin := model.User{
			ExternalID: "superadmin-001",
			Email:      "muttaqinfaiz@gmail.com",
			Username:   "superadmin",
			FirstName:  "Muttaqin",
			LastName:   "Faiz",
			RoleID:     1, // Super Admin
			Status:     model.StatusActive,
		}
		if err := db.Create(&superAdmin).Error; err != nil {
			return fmt.Errorf("failed creating super admin: %w", err)
		}

		// Default User 1
		defaultUser1 := model.User{
			ExternalID: "user-001",
			Email:      "faizipb@gmail.com",
			Username:   "faizipb",
			FirstName:  "Faiz",
			LastName:   "IPB",
			RoleID:     2, // Default/Customer
			Status:     model.StatusActive,
		}
		if err := db.Create(&defaultUser1).Error; err != nil {
			return fmt.Errorf("failed creating default user 1: %w", err)
		}

		// Seller User with Shop
		sellerUser := model.User{
			ExternalID: "seller-001",
			Email:      "pojok.brn@gmail.com",
			Username:   "pojokbrn",
			FirstName:  "Pojok",
			LastName:   "BRN",
			RoleID:     3, // Seller
			Status:     model.StatusActive,
		}
		if err := db.Create(&sellerUser).Error; err != nil {
			return fmt.Errorf("failed creating seller user: %w", err)
		}

		// Create shop for seller
		sellerShop := model.Shop{
			UserID:      sellerUser.ID,
			Name:        "Pojok BRN Store",
			Slug:        "pojok-brn-store",
			City:        "Bandung",
			Description: "Official Pojok BRN Store - Quality products for everyone",
			IsOfficial:  false,
			IsActive:    true,
		}
		if err := db.Create(&sellerShop).Error; err != nil {
			return fmt.Errorf("failed creating seller shop: %w", err)
		}

		// Create super admin shop (optional - admins can have shops too)
		superAdminShop := model.Shop{
			UserID:      superAdmin.ID,
			Name:        "Admin Official Store",
			Slug:        "admin-official-store",
			City:        "Jakarta",
			Description: "Official Admin Store - Premium products",
			IsOfficial:  true,
			IsActive:    true,
		}
		if err := db.Create(&superAdminShop).Error; err != nil {
			return fmt.Errorf("failed creating admin shop: %w", err)
		}
	}

	return nil
}
