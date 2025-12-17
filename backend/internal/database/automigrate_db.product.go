package database

import (
	"fmt"
	"math/rand"

	"github.com/faiz-muttaqin/lgs/backend/internal/model"
	"github.com/faiz-muttaqin/lgs/backend/pkg/types"
	"gorm.io/gorm"
)

func AutoMigrateDBProduct(db *gorm.DB) error {
	// Auto migrate all product-related tables
	if err := db.AutoMigrate(
		&model.Category{},
		&model.SubCategory{},
		&model.Shop{},
		&model.Product{},
		&model.ProductImage{},
		&model.ProductLabel{},
		&model.ProductBadge{},
		&model.ProductVariant{},
	); err != nil {
		return err
	}

	// Check if data already exists
	var categoryCount int64
	db.Model(&model.Category{}).Count(&categoryCount)
	if categoryCount > 0 {
		return nil // Data already seeded
	}

	// Image URLs from media_image_urls.txt
	imageURLs := []string{
		"https://images.unsplash.com/photo-1602143407151-7111542de6e8?w=400&auto=format&fit=crop&q=60&ixlib=rb-4.1.0&ixid=M3wxMjA3fDB8MHxzZWFyY2h8NHx8cHJvZHVjdHxlbnwwfHwwfHx8MA%3D%3D",
		"https://images.unsplash.com/photo-1615397349754-cfa2066a298e?w=400&auto=format&fit=crop&q=60&ixlib=rb-4.1.0&ixid=M3wxMjA3fDB8MHxzZWFyY2h8MTF8fHByb2R1Y3R8ZW58MHx8MHx8fDA%3D",
		"https://images.unsplash.com/photo-1543163521-1bf539c55dd2?w=400&auto=format&fit=crop&q=60&ixlib=rb-4.1.0&ixid=M3wxMjA3fDB8MHxzZWFyY2h8MTV8fHByb2R1Y3R8ZW58MHx8MHx8fDA%3D",
		"https://images.unsplash.com/photo-1549049950-48d5887197a0?w=400&auto=format&fit=crop&q=60&ixlib=rb-4.1.0&ixid=M3wxMjA3fDB8MHxzZWFyY2h8MjB8fHByb2R1Y3R8ZW58MHx8MHx8fDA%3D",
		"https://images.unsplash.com/photo-1523275335684-37898b6baf30?w=400&auto=format&fit=crop&q=60&ixlib=rb-4.1.0&ixid=M3wxMjA3fDB8MHxzZWFyY2h8Mnx8cHJvZHVjdHxlbnwwfHwwfHx8MA%3D%3D",
		"https://images.unsplash.com/photo-1505740420928-5e560c06d30e?w=400&auto=format&fit=crop&q=60&ixlib=rb-4.1.0&ixid=M3wxMjA3fDB8MHxzZWFyY2h8M3x8cHJvZHVjdHxlbnwwfHwwfHx8MA%3D%3D",
		"https://images.unsplash.com/photo-1541643600914-78b084683601?w=400&auto=format&fit=crop&q=60&ixlib=rb-4.1.0&ixid=M3wxMjA3fDB8MHxzZWFyY2h8N3x8cHJvZHVjdHxlbnwwfHwwfHx8MA%3D%3D",
		"https://images.unsplash.com/photo-1546868871-7041f2a55e12?w=400&auto=format&fit=crop&q=60&ixlib=rb-4.1.0&ixid=M3wxMjA3fDB8MHxzZWFyY2h8MTJ8fHByb2R1Y3R8ZW58MHx8MHx8fDA%3D",
		"https://images.unsplash.com/photo-1615396899839-c99c121888b0?w=400&auto=format&fit=crop&q=60&ixlib=rb-4.1.0&ixid=M3wxMjA3fDB8MHxzZWFyY2h8MTR8fHByb2R1Y3R8ZW58MHx8MHx8fDA%3D",
		"https://images.unsplash.com/photo-1560343090-f0409e92791a?w=400&auto=format&fit=crop&q=60&ixlib=rb-4.1.0&ixid=M3wxMjA3fDB8MHxzZWFyY2h8MTl8fHByb2R1Y3R8ZW58MHx8MHx8fDA%3D",
		"https://images.unsplash.com/photo-1526170375885-4d8ecf77b99f?w=400&auto=format&fit=crop&q=60&ixlib=rb-4.1.0&ixid=M3wxMjA3fDB8MHxzZWFyY2h8Nnx8cHJvZHVjdHxlbnwwfHwwfHx8MA%3D%3D",
		"https://images.unsplash.com/photo-1542291026-7eec264c27ff?w=400&auto=format&fit=crop&q=60&ixlib=rb-4.1.0&ixid=M3wxMjA3fDB8MHxzZWFyY2h8OHx8cHJvZHVjdHxlbnwwfHwwfHx8MA%3D%3D",
		"https://images.unsplash.com/photo-1572635196237-14b3f281503f?w=400&auto=format&fit=crop&q=60&ixlib=rb-4.1.0&ixid=M3wxMjA3fDB8MHxzZWFyY2h8MTB8fHByb2R1Y3R8ZW58MHx8MHx8fDA%3D",
		"https://images.unsplash.com/photo-1491553895911-0055eca6402d?w=400&auto=format&fit=crop&q=60&ixlib=rb-4.1.0&ixid=M3wxMjA3fDB8MHxzZWFyY2h8MTZ8fHByb2R1Y3R8ZW58MHx8MHx8fDA%3D",
		"https://images.unsplash.com/photo-1503602642458-232111445657?w=400&auto=format&fit=crop&q=60&ixlib=rb-4.1.0&ixid=M3wxMjA3fDB8MHxzZWFyY2h8MTh8fHByb2R1Y3R8ZW58MHx8MHx8fDA%3D",
		"https://images.unsplash.com/photo-1586495777744-4413f21062fa?w=400&auto=format&fit=crop&q=60&ixlib=rb-4.1.0&ixid=M3wxMjA3fDB8MHxzZWFyY2h8MjN8fHByb2R1Y3R8ZW58MHx8MHx8fDA%3D",
		"https://images.unsplash.com/photo-1581235720704-06d3acfcb36f?w=400&auto=format&fit=crop&q=60&ixlib=rb-4.1.0&ixid=M3wxMjA3fDB8MHxzZWFyY2h8Mjh8fHByb2R1Y3R8ZW58MHx8MHx8fDA%3D",
		"https://images.unsplash.com/photo-1620987278429-ab178d6eb547?w=400&auto=format&fit=crop&q=60&ixlib=rb-4.1.0&ixid=M3wxMjA3fDB8MHxzZWFyY2h8MzB8fHByb2R1Y3R8ZW58MHx8MHx8fDA%3D",
		"https://images.unsplash.com/photo-1524638067-feba7e8ed70f?w=400&auto=format&fit=crop&q=60&ixlib=rb-4.1.0&ixid=M3wxMjA3fDB8MHxzZWFyY2h8MzZ8fHByb2R1Y3R8ZW58MHx8MHx8fDA%3D",
		"https://images.unsplash.com/photo-1491637639811-60e2756cc1c7?w=400&auto=format&fit=crop&q=60&ixlib=rb-4.1.0&ixid=M3wxMjA3fDB8MHxzZWFyY2h8MjR8fHByb2R1Y3R8ZW58MHx8MHx8fDA%3D",
		"https://images.unsplash.com/photo-1504274066651-8d31a536b11a?w=400&auto=format&fit=crop&q=60&ixlib=rb-4.1.0&ixid=M3wxMjA3fDB8MHxzZWFyY2h8MjZ8fHByb2R1Y3R8ZW58MHx8MHx8fDA%3D",
		"https://images.unsplash.com/photo-1556227834-09f1de7a7d14?w=400&auto=format&fit=crop&q=60&ixlib=rb-4.1.0&ixid=M3wxMjA3fDB8MHxzZWFyY2h8MzF8fHByb2R1Y3R8ZW58MHx8MHx8fDA%3D",
		"https://images.unsplash.com/photo-1583394838336-acd977736f90?w=400&auto=format&fit=crop&q=60&ixlib=rb-4.1.0&ixid=M3wxMjA3fDB8MHxzZWFyY2h8MjJ8fHByb2R1Y3R8ZW58MHx8MHx8fDA%3D",
		"https://images.unsplash.com/photo-1553456558-aff63285bdd1?w=400&auto=format&fit=crop&q=60&ixlib=rb-4.1.0&ixid=M3wxMjA3fDB8MHxzZWFyY2h8Mjd8fHByb2R1Y3R8ZW58MHx8MHx8fDA%3D",
		"https://images.unsplash.com/photo-1560769629-975ec94e6a86?w=400&auto=format&fit=crop&q=60&ixlib=rb-4.1.0&ixid=M3wxMjA3fDB8MHxzZWFyY2h8MzJ8fHByb2R1Y3R8ZW58MHx8MHx8fDA%3D",
		"https://images.unsplash.com/photo-1564466809058-bf4114d55352?w=400&auto=format&fit=crop&q=60&ixlib=rb-4.1.0&ixid=M3wxMjA3fDB8MHxzZWFyY2h8MzV8fHByb2R1Y3R8ZW58MHx8MHx8fDA%3D",
		"https://images.unsplash.com/photo-1580870069867-74c57ee1bb07?w=400&auto=format&fit=crop&q=60&ixlib=rb-4.1.0&ixid=M3wxMjA3fDB8MHxzZWFyY2h8MzR8fHByb2R1Y3R8ZW58MHx8MHx8fDA%3D",
		"https://images.unsplash.com/photo-1556228578-8c89e6adf883?w=400&auto=format&fit=crop&q=60&ixlib=rb-4.1.0&ixid=M3wxMjA3fDB8MHxzZWFyY2h8Mzh8fHByb2R1Y3R8ZW58MHx8MHx8fDA%3D",
		"https://images.unsplash.com/photo-1532298229144-0ec0c57515c7?w=400&auto=format&fit=crop&q=60&ixlib=rb-4.1.0&ixid=M3wxMjA3fDB8MHxzZWFyY2h8Mzl8fHByb2R1Y3R8ZW58MHx8MHx8fDA%3D",
		"https://images.unsplash.com/photo-1485955900006-10f4d324d411?w=400&auto=format&fit=crop&q=60&ixlib=rb-4.1.0&ixid=M3wxMjA3fDB8MHxzZWFyY2h8NDB8fHByb2R1Y3R8ZW58MHx8MHx8fDA%3D",
		"https://images.unsplash.com/photo-1556228578-567ba127e37f?w=400&auto=format&fit=crop&q=60&ixlib=rb-4.1.0&ixid=M3wxMjA3fDB8MHxzZWFyY2h8NDN8fHByb2R1Y3R8ZW58MHx8MHx8fDA%3D",
		"https://images.unsplash.com/photo-1519669011783-4eaa95fa1b7d?w=400&auto=format&fit=crop&q=60&ixlib=rb-4.1.0&ixid=M3wxMjA3fDB8MHxzZWFyY2h8NDh8fHByb2R1Y3R8ZW58MHx8MHx8fDA%3D",
		"https://images.unsplash.com/photo-1553062407-98eeb64c6a62?w=400&auto=format&fit=crop&q=60&ixlib=rb-4.1.0&ixid=M3wxMjA3fDB8MHxzZWFyY2h8NTF8fHByb2R1Y3R8ZW58MHx8MHx8fDA%3D",
		"https://images.unsplash.com/photo-1547949003-9792a18a2601?w=400&auto=format&fit=crop&q=60&ixlib=rb-4.1.0&ixid=M3wxMjA3fDB8MHxzZWFyY2h8NDR8fHByb2R1Y3R8ZW58MHx8MHx8fDA%3D",
		"https://images.unsplash.com/photo-1511499767150-a48a237f0083?w=400&auto=format&fit=crop&q=60&ixlib=rb-4.1.0&ixid=M3wxMjA3fDB8MHxzZWFyY2h8NDZ8fHByb2R1Y3R8ZW58MHx8MHx8fDA%3D",
		"https://images.unsplash.com/photo-1547887537-6158d64c35b3?w=400&auto=format&fit=crop&q=60&ixlib=rb-4.1.0&ixid=M3wxMjA3fDB8MHxzZWFyY2h8NDd8fHByb2R1Y3R8ZW58MHx8MHx8fDA%3D",
		"https://images.unsplash.com/photo-1556228578-0d85b1a4d571?w=400&auto=format&fit=crop&q=60&ixlib=rb-4.1.0&ixid=M3wxMjA3fDB8MHxzZWFyY2h8NDJ8fHByb2R1Y3R8ZW58MHx8MHx8fDA%3D",
		"https://images.unsplash.com/photo-1525904097878-94fb15835963?w=400&auto=format&fit=crop&q=60&ixlib=rb-4.1.0&ixid=M3wxMjA3fDB8MHxzZWFyY2h8NTZ8fHByb2R1Y3R8ZW58MHx8MHx8fDA%3D",
		"https://images.unsplash.com/photo-1525966222134-fcfa99b8ae77?w=400&auto=format&fit=crop&q=60&ixlib=rb-4.1.0&ixid=M3wxMjA3fDB8MHxzZWFyY2h8NTV8fHByb2R1Y3R8ZW58MHx8MHx8fDA%3D",
		"https://images.unsplash.com/photo-1509695507497-903c140c43b0?w=400&auto=format&fit=crop&q=60&ixlib=rb-4.1.0&ixid=M3wxMjA3fDB8MHxzZWFyY2h8NTB8fHByb2R1Y3R8ZW58MHx8MHx8fDA%3D",
		"https://images.unsplash.com/photo-1620916566398-39f1143ab7be?w=400&auto=format&fit=crop&q=60&ixlib=rb-4.1.0&ixid=M3wxMjA3fDB8MHxzZWFyY2h8NTl8fHByb2R1Y3R8ZW58MHx8MHx8fDA%3D",
		"https://images.unsplash.com/photo-1486401899868-0e435ed85128?w=400&auto=format&fit=crop&q=60&ixlib=rb-4.1.0&ixid=M3wxMjA3fDB8MHxzZWFyY2h8NTJ8fHByb2R1Y3R8ZW58MHx8MHx8fDA%3D",
		"https://images.unsplash.com/photo-1567721913486-6585f069b332?w=400&auto=format&fit=crop&q=60&ixlib=rb-4.1.0&ixid=M3wxMjA3fDB8MHxzZWFyY2h8NTR8fHByb2R1Y3R8ZW58MHx8MHx8fDA%3D",
		"https://images.unsplash.com/photo-1543512214-318c7553f230?w=400&auto=format&fit=crop&q=60&ixlib=rb-4.1.0&ixid=M3wxMjA3fDB8MHxzZWFyY2h8NTh8fHByb2R1Y3R8ZW58MHx8MHx8fDA%3D",
		"https://images.unsplash.com/photo-1571781926291-c477ebfd024b?w=400&auto=format&fit=crop&q=60&ixlib=rb-4.1.0&ixid=M3wxMjA3fDB8MHxzZWFyY2h8NjJ8fHByb2R1Y3R8ZW58MHx8MHx8fDA%3D",
		"https://images.unsplash.com/photo-1608571423902-eed4a5ad8108?w=400&auto=format&fit=crop&q=60&ixlib=rb-4.1.0&ixid=M3wxMjA3fDB8MHxzZWFyY2h8NjB8fHByb2R1Y3R8ZW58MHx8MHx8fDA%3D",
		"https://images.unsplash.com/photo-1570831739435-6601aa3fa4fb?w=400&auto=format&fit=crop&q=60&ixlib=rb-4.1.0&ixid=M3wxMjA3fDB8MHxzZWFyY2h8NjR8fHByb2R1Y3R8ZW58MHx8MHx8fDA%3D",
		"https://images.unsplash.com/photo-1545127398-14699f92334b?w=400&auto=format&fit=crop&q=60&ixlib=rb-4.1.0&ixid=M3wxMjA3fDB8MHxzZWFyY2h8Njd8fHByb2R1Y3R8ZW58MHx8MHx8fDA%3D",
		"https://images.unsplash.com/photo-1611930021592-a8cfd5319ceb?w=400&auto=format&fit=crop&q=60&ixlib=rb-4.1.0&ixid=M3wxMjA3fDB8MHxzZWFyY2h8NzB8fHByb2R1Y3R8ZW58MHx8MHx8fDA%3D",
		"https://images.unsplash.com/photo-1608248543803-ba4f8c70ae0b?w=400&auto=format&fit=crop&q=60&ixlib=rb-4.1.0&ixid=M3wxMjA3fDB8MHxzZWFyY2h8NjN8fHByb2R1Y3R8ZW58MHx8MHx8fDA%3D",
		"https://images.unsplash.com/photo-1559056199-641a0ac8b55e?w=400&auto=format&fit=crop&q=60&ixlib=rb-4.1.0&ixid=M3wxMjA3fDB8MHxzZWFyY2h8NjZ8fHByb2R1Y3R8ZW58MHx8MHx8fDA%3D",
		"https://images.unsplash.com/photo-1524805444758-089113d48a6d?w=400&auto=format&fit=crop&q=60&ixlib=rb-4.1.0&ixid=M3wxMjA3fDB8MHxzZWFyY2h8Njh8fHByb2R1Y3R8ZW58MHx8MHx8fDA%3D",
		"https://images.unsplash.com/photo-1530630458144-014709e10016?w=400&auto=format&fit=crop&q=60&ixlib=rb-4.1.0&ixid=M3wxMjA3fDB8MHxzZWFyY2h8NzF8fHByb2R1Y3R8ZW58MHx8MHx8fDA%3D",
		"https://images.unsplash.com/photo-1563170351-be82bc888aa4?w=400&auto=format&fit=crop&q=60&ixlib=rb-4.1.0&ixid=M3wxMjA3fDB8MHxzZWFyY2h8NzJ8fHByb2R1Y3R8ZW58MHx8MHx8fDA%3D",
		"https://images.unsplash.com/photo-1597317628840-d3472f7aa7fc?w=400&auto=format&fit=crop&q=60&ixlib=rb-4.1.0&ixid=M3wxMjA3fDB8MHxzZWFyY2h8NzZ8fHByb2R1Y3R8ZW58MHx8MHx8fDA%3D",
		"https://images.unsplash.com/photo-1556228578-f9707385e031?w=400&auto=format&fit=crop&q=60&ixlib=rb-4.1.0&ixid=M3wxMjA3fDB8MHxzZWFyY2h8Nzl8fHByb2R1Y3R8ZW58MHx8MHx8fDA%3D",
		"https://images.unsplash.com/photo-1511556820780-d912e42b4980?w=400&auto=format&fit=crop&q=60&ixlib=rb-4.1.0&ixid=M3wxMjA3fDB8MHxzZWFyY2h8NzV8fHByb2R1Y3R8ZW58MHx8MHx8fDA%3D",
		"https://images.unsplash.com/photo-1549482199-bc1ca6f58502?w=400&auto=format&fit=crop&q=60&ixlib=rb-4.1.0&ixid=M3wxMjA3fDB8MHxzZWFyY2h8ODB8fHByb2R1Y3R8ZW58MHx8MHx8fDA%3D",
		"https://images.unsplash.com/photo-1560393464-5c69a73c5770?w=400&auto=format&fit=crop&q=60&ixlib=rb-4.1.0&ixid=M3wxMjA3fDB8MHxzZWFyY2h8ODJ8fHByb2R1Y3R8ZW58MHx8MHx8fDA%3D",
		"https://images.unsplash.com/photo-1608528577891-eb055944f2e7?w=400&auto=format&fit=crop&q=60&ixlib=rb-4.1.0&ixid=M3wxMjA3fDB8MHxzZWFyY2h8Nzh8fHByb2R1Y3R8ZW58MHx8MHx8fDA%3D",
		"https://images.unsplash.com/photo-1526429257838-9bf73dd45097?w=400&auto=format&fit=crop&q=60&ixlib=rb-4.1.0&ixid=M3wxMjA3fDB8MHxzZWFyY2h8ODN8fHByb2R1Y3R8ZW58MHx8MHx8fDA%3D",
		"https://images.unsplash.com/photo-1524678606370-a47ad25cb82a?w=400&auto=format&fit=crop&q=60&ixlib=rb-4.1.0&ixid=M3wxMjA3fDB8MHxzZWFyY2h8ODZ8fHByb2R1Y3R8ZW58MHx8MHx8fDA%3D",
		"https://images.unsplash.com/photo-1530914547840-346c183410de?w=400&auto=format&fit=crop&q=60&ixlib=rb-4.1.0&ixid=M3wxMjA3fDB8MHxzZWFyY2h8OTF8fHByb2R1Y3R8ZW58MHx8MHx8fDA%3D",
		"https://images.unsplash.com/photo-1526947425960-945c6e72858f?w=400&auto=format&fit=crop&q=60&ixlib=rb-4.1.0&ixid=M3wxMjA3fDB8MHxzZWFyY2h8OTR8fHByb2R1Y3R8ZW58MHx8MHx8fDA%3D",
		"https://images.unsplash.com/photo-1479064555552-3ef4979f8908?w=400&auto=format&fit=crop&q=60&ixlib=rb-4.1.0&ixid=M3wxMjA3fDB8MHxzZWFyY2h8ODR8fHByb2R1Y3R8ZW58MHx8MHx8fDA%3D",
		"https://images.unsplash.com/photo-1611930022073-b7a4ba5fcccd?w=400&auto=format&fit=crop&q=60&ixlib=rb-4.1.0&ixid=M3wxMjA3fDB8MHxzZWFyY2h8ODd8fHByb2R1Y3R8ZW58MHx8MHx8fDA%3D",
		"https://images.unsplash.com/photo-1527864550417-7fd91fc51a46?w=400&auto=format&fit=crop&q=60&ixlib=rb-4.1.0&ixid=M3wxMjA3fDB8MHxzZWFyY2h8ODh8fHByb2R1Y3R8ZW58MHx8MHx8fDA%3D",
		"https://images.unsplash.com/photo-1583947215259-38e31be8751f?w=400&auto=format&fit=crop&q=60&ixlib=rb-4.1.0&ixid=M3wxMjA3fDB8MHxzZWFyY2h8OTB8fHByb2R1Y3R8ZW58MHx8MHx8fDA%3D",
		"https://images.unsplash.com/photo-1535585209827-a15fcdbc4c2d?w=400&auto=format&fit=crop&q=60&ixlib=rb-4.1.0&ixid=M3wxMjA3fDB8MHxzZWFyY2h8OTV8fHByb2R1Y3R8ZW58MHx8MHx8fDA%3D",
		"https://images.unsplash.com/photo-1601049541289-9b1b7bbbfe19?w=400&auto=format&fit=crop&q=60&ixlib=rb-4.1.0&ixid=M3wxMjA3fDB8MHxzZWFyY2h8OTl8fHByb2R1Y3R8ZW58MHx8MHx8fDA%3D",
		"https://images.unsplash.com/photo-1555487505-8603a1a69755?w=400&auto=format&fit=crop&q=60&ixlib=rb-4.1.0&ixid=M3wxMjA3fDB8MHxzZWFyY2h8OTZ8fHByb2R1Y3R8ZW58MHx8MHx8fDA%3D",
		"https://images.unsplash.com/photo-1556228720-195a672e8a03?w=400&auto=format&fit=crop&q=60&ixlib=rb-4.1.0&ixid=M3wxMjA3fDB8MHxzZWFyY2h8OTJ8fHByb2R1Y3R8ZW58MHx8MHx8fDA%3D",
		"https://images.unsplash.com/photo-1522115174737-2497162f69ec?w=400&auto=format&fit=crop&q=60&ixlib=rb-4.1.0&ixid=M3wxMjA3fDB8MHxzZWFyY2h8OTh8fHByb2R1Y3R8ZW58MHx8MHx8fDA%3D",
		"https://images.unsplash.com/photo-1532667449560-72a95c8d381b?fm=jpg&q=60&w=3000&ixlib=rb-4.1.0&ixid=M3wxMjA3fDB8MHxzZWFyY2h8MTA0fHxwcm9kdWN0fGVufDB8fDB8fHww",
		"https://images.unsplash.com/photo-1522643628976-0a170f6722ab?fm=jpg&q=60&w=3000&ixlib=rb-4.1.0&ixid=M3wxMjA3fDB8MHxzZWFyY2h8MTA2fHxwcm9kdWN0fGVufDB8fDB8fHww",
		"https://images.unsplash.com/photo-1629198688000-71f23e745b6e?fm=jpg&q=60&w=3000&ixlib=rb-4.1.0&ixid=M3wxMjA3fDB8MHxzZWFyY2h8MTEwfHxwcm9kdWN0fGVufDB8fDB8fHww",
	}

	// Seed Categories
	categories := []model.Category{
		{Name: "Electronics", Slug: "electronics", Description: "Electronic devices and accessories", ImageURL: imageURLs[0], IsActive: true},
		{Name: "Fashion", Slug: "fashion", Description: "Clothing, shoes, and accessories", ImageURL: imageURLs[1], IsActive: true},
		{Name: "Home & Kitchen", Slug: "home-kitchen", Description: "Home appliances and kitchen tools", ImageURL: imageURLs[2], IsActive: true},
		{Name: "Sports & Outdoors", Slug: "sports-outdoors", Description: "Sports equipment and outdoor gear", ImageURL: imageURLs[3], IsActive: true},
		{Name: "Beauty & Health", Slug: "beauty-health", Description: "Beauty products and health care", ImageURL: imageURLs[4], IsActive: true},
	}
	if err := db.Create(&categories).Error; err != nil {
		return fmt.Errorf("failed creating categories: %w", err)
	}

	// Seed SubCategories
	subCategories := []model.SubCategory{
		{CategoryID: categories[0].ID, Name: "Smartphones", Slug: "smartphones", Description: "Mobile phones and accessories", ImageURL: imageURLs[5], IsActive: true},
		{CategoryID: categories[0].ID, Name: "Laptops", Slug: "laptops", Description: "Notebooks and laptop accessories", ImageURL: imageURLs[6], IsActive: true},
		{CategoryID: categories[0].ID, Name: "Headphones", Slug: "headphones", Description: "Audio devices", ImageURL: imageURLs[7], IsActive: true},
		{CategoryID: categories[1].ID, Name: "Men's Clothing", Slug: "mens-clothing", Description: "Men's fashion items", ImageURL: imageURLs[8], IsActive: true},
		{CategoryID: categories[1].ID, Name: "Women's Clothing", Slug: "womens-clothing", Description: "Women's fashion items", ImageURL: imageURLs[9], IsActive: true},
		{CategoryID: categories[1].ID, Name: "Shoes", Slug: "shoes", Description: "Footwear for all", ImageURL: imageURLs[10], IsActive: true},
		{CategoryID: categories[2].ID, Name: "Coffee Makers", Slug: "coffee-makers", Description: "Coffee and tea makers", ImageURL: imageURLs[11], IsActive: true},
		{CategoryID: categories[2].ID, Name: "Cookware", Slug: "cookware", Description: "Cooking tools and utensils", ImageURL: imageURLs[12], IsActive: true},
		{CategoryID: categories[3].ID, Name: "Fitness Equipment", Slug: "fitness-equipment", Description: "Gym and fitness gear", ImageURL: imageURLs[13], IsActive: true},
		{CategoryID: categories[3].ID, Name: "Camping Gear", Slug: "camping-gear", Description: "Outdoor camping equipment", ImageURL: imageURLs[14], IsActive: true},
		{CategoryID: categories[4].ID, Name: "Skincare", Slug: "skincare", Description: "Skincare products", ImageURL: imageURLs[15], IsActive: true},
		{CategoryID: categories[4].ID, Name: "Makeup", Slug: "makeup", Description: "Cosmetics and makeup", ImageURL: imageURLs[16], IsActive: true},
	}
	if err := db.Create(&subCategories).Error; err != nil {
		return fmt.Errorf("failed creating subcategories: %w", err)
	}

	// Get existing shops (created by users) and create additional demo shops if needed
	var existingShops []model.Shop
	db.Find(&existingShops)

	// If no shops exist or we need more, create additional demo shops
	var shops []model.Shop
	if len(existingShops) == 0 {
		// Create demo shops only if none exist
		// Note: In production, these would be created by actual users
		shops = []model.Shop{
			{UserID: 999, Name: "TechHub Official", Slug: "techhub-official", Domain: "techhub.com", City: "Jakarta", ImageURL: imageURLs[17], IsOfficial: true, IsActive: true},
			{UserID: 998, Name: "Fashion Store", Slug: "fashion-store", Domain: "fashionstore.com", City: "Bandung", ImageURL: imageURLs[18], IsOfficial: true, IsActive: true},
			{UserID: 997, Name: "HomeGoods Market", Slug: "homegoods-market", Domain: "homegoods.com", City: "Surabaya", ImageURL: imageURLs[19], IsOfficial: false, IsActive: true},
			{UserID: 996, Name: "Sports Pro", Slug: "sports-pro", Domain: "sportspro.com", City: "Medan", ImageURL: imageURLs[20], IsOfficial: true, IsActive: true},
			{UserID: 995, Name: "Beauty Plus", Slug: "beauty-plus", Domain: "beautyplus.com", City: "Semarang", ImageURL: imageURLs[21], IsOfficial: false, IsActive: true},
			{UserID: 994, Name: "Gadget World", Slug: "gadget-world", Domain: "gadgetworld.com", City: "Yogyakarta", ImageURL: imageURLs[22], IsOfficial: true, IsActive: true},
			{UserID: 993, Name: "Style Avenue", Slug: "style-avenue", Domain: "styleavenue.com", City: "Malang", ImageURL: imageURLs[23], IsOfficial: false, IsActive: true},
			{UserID: 992, Name: "Kitchen Master", Slug: "kitchen-master", Domain: "kitchenmaster.com", City: "Palembang", ImageURL: imageURLs[24], IsOfficial: true, IsActive: true},
		}
		// Create these demo shops without UserID constraint (for demo purposes only)
		for i := range shops {
			db.Exec("INSERT INTO shops (user_id, name, slug, domain, city, image_url, is_official, is_active, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?, ?, ?, NOW(), NOW())",
				shops[i].UserID, shops[i].Name, shops[i].Slug, shops[i].Domain, shops[i].City, shops[i].ImageURL, shops[i].IsOfficial, shops[i].IsActive)
		}
		db.Find(&shops)
	} else {
		shops = existingShops
	}

	// Product names by category
	productNames := map[uint][]string{
		subCategories[0].ID: { // Smartphones
			"iPhone 15 Pro Max", "Samsung Galaxy S24 Ultra", "Google Pixel 8 Pro", "OnePlus 12", "Xiaomi 14 Pro",
			"OPPO Find X7 Ultra", "Vivo X100 Pro", "Realme GT 5 Pro", "Nothing Phone 2", "Sony Xperia 5 V",
		},
		subCategories[1].ID: { // Laptops
			"MacBook Pro M3", "Dell XPS 15", "HP Spectre x360", "Lenovo ThinkPad X1", "ASUS ROG Zephyrus",
			"Acer Swift 3", "Microsoft Surface Laptop", "Razer Blade 15", "MSI Prestige 14", "LG Gram 17",
		},
		subCategories[2].ID: { // Headphones
			"Sony WH-1000XM5", "Bose QuietComfort Ultra", "AirPods Pro 2", "Sennheiser Momentum 4", "JBL Tune 760NC",
			"Beats Studio Pro", "Audio-Technica ATH-M50x", "Anker Soundcore Q45", "Skullcandy Crusher Evo", "Jabra Elite 85h",
		},
		subCategories[3].ID: { // Men's Clothing
			"Classic Polo Shirt", "Denim Jacket", "Formal Dress Shirt", "Casual Chinos", "Wool Sweater",
			"Leather Jacket", "Cotton T-Shirt Pack", "Slim Fit Jeans", "Sport Hoodie", "Blazer Suit",
		},
		subCategories[4].ID: { // Women's Clothing
			"Floral Summer Dress", "Elegant Evening Gown", "Casual Blouse", "High-Waist Jeans", "Cardigan Sweater",
			"Maxi Dress", "Business Suit", "Yoga Pants", "Leather Skirt", "Kimono Jacket",
		},
		subCategories[5].ID: { // Shoes
			"Running Sneakers", "Leather Boots", "Canvas Slip-Ons", "High Heels", "Hiking Boots",
			"Sandals", "Loafers", "Basketball Shoes", "Ballet Flats", "Ankle Boots",
		},
		subCategories[6].ID: { // Coffee Makers
			"Espresso Machine Pro", "Drip Coffee Maker", "French Press Premium", "Pour Over Set", "Coffee Grinder Electric",
			"Turkish Coffee Maker", "Cold Brew Pitcher", "Moka Pot Classic", "Single Serve Coffee Maker", "Percolator Stainless",
		},
		subCategories[7].ID: { // Cookware
			"Non-Stick Pan Set", "Cast Iron Skillet", "Stainless Steel Pot", "Wok Carbon Steel", "Baking Sheet Set",
			"Dutch Oven", "Pressure Cooker", "Frying Pan Ceramic", "Sauce Pan Set", "Grill Pan",
		},
		subCategories[8].ID: { // Fitness Equipment
			"Adjustable Dumbbells", "Yoga Mat Premium", "Resistance Bands Set", "Kettlebell Set", "Treadmill Foldable",
			"Exercise Bike", "Pull-Up Bar", "Ab Roller", "Jump Rope", "Weight Bench",
		},
		subCategories[9].ID: { // Camping Gear
			"Tent 4-Person", "Sleeping Bag Warm", "Camping Stove Portable", "Backpack Hiking", "Camping Chair Foldable",
			"Cooler Box", "Lantern LED", "Hammock Double", "Water Filter Portable", "Multi-Tool Camping",
		},
		subCategories[10].ID: { // Skincare
			"Vitamin C Serum", "Hyaluronic Acid Moisturizer", "Retinol Night Cream", "Sunscreen SPF 50", "Cleanser Gentle",
			"Face Mask Sheet", "Eye Cream Anti-Aging", "Toner Alcohol-Free", "Exfoliating Scrub", "Lip Balm Hydrating",
		},
		subCategories[11].ID: { // Makeup
			"Foundation Matte", "Eyeshadow Palette", "Mascara Volumizing", "Lipstick Set", "Blush Powder",
			"Highlighter Glow", "Eyeliner Waterproof", "Bronzer Contour", "Setting Spray", "Makeup Brush Set",
		},
	}

	products := []model.Product{}
	imgIdx := 25

	// Generate 100+ products
	for subCatID, names := range productNames {
		for _, name := range names {
			price := float64(rand.Intn(500000) + 50000)
			slashedPrice := price * 1.15
			discountPct := 5 + rand.Intn(45)

			product := model.Product{
				SKU:           fmt.Sprintf("SKU-%d-%d", subCatID, len(products)+1),
				Name:          name,
				Slug:          fmt.Sprintf("%s-%d", slugify(name), len(products)+1),
				Subtitle:      "Premium quality product",
				Description:   fmt.Sprintf("High-quality %s with excellent features and performance. Perfect for daily use.", name),
				ImageURL:      imageURLs[imgIdx%len(imageURLs)],
				Price:         price,
				SlashedPrice:  slashedPrice,
				DiscountPct:   discountPct,
				Stock:         rand.Intn(500) + 10,
				Rating:        4.0 + rand.Float32(),
				CountReview:   rand.Intn(5000) + 100,
				CountSold:     rand.Intn(10000) + 500,
				Weight:        rand.Intn(5000) + 100,
				IsActive:      true,
				IsFeatured:    rand.Intn(10) < 3,
				Status:        types.Badge(model.ProductStatusPublished),
				SubCategoryID: subCatID,
				ShopID:        shops[rand.Intn(len(shops))].ID,
			}

			// Set CategoryID based on SubCategory
			for _, subCat := range subCategories {
				if subCat.ID == subCatID {
					product.CategoryID = subCat.CategoryID
					break
				}
			}

			products = append(products, product)
			imgIdx++
		}
	}

	// Create products in batches
	batchSize := 50
	for i := 0; i < len(products); i += batchSize {
		end := i + batchSize
		if end > len(products) {
			end = len(products)
		}
		if err := db.Create(products[i:end]).Error; err != nil {
			return fmt.Errorf("failed creating products batch: %w", err)
		}
	}

	// Add additional product images, labels, and badges for some products
	for i, product := range products {
		if i%3 == 0 { // Add extra images to every 3rd product
			additionalImages := []model.ProductImage{
				{ProductID: product.ID, ImageURL: imageURLs[(imgIdx+i)%len(imageURLs)], Position: 1, IsMain: false},
				{ProductID: product.ID, ImageURL: imageURLs[(imgIdx+i+1)%len(imageURLs)], Position: 2, IsMain: false},
			}
			db.Create(&additionalImages)
		}

		if i%5 == 0 { // Add labels to every 5th product
			labels := []model.ProductLabel{
				{ProductID: product.ID, Title: "Flash Sale", Type: "red", Position: "overlay_1"},
				{ProductID: product.ID, Title: "Best Seller", Type: "blue", Position: "campaign"},
			}
			db.Create(&labels)
		}

		if i%7 == 0 { // Add badges to every 7th product
			badges := []model.ProductBadge{
				{ProductID: product.ID, Title: "Official Store", ImageURL: imageURLs[(imgIdx+i+2)%len(imageURLs)]},
			}
			db.Create(&badges)
		}
	}

	return nil
}
