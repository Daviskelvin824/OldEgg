package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	ProductID          string  `json:"product_id"`
	ShopID             int     `json:"shop_id" gorm:"foreignKey:ID"`
	ProductCategoryID  int     `json:"product_category_id" gorm:"foreignKey:ProductCategoryID"`
	ProductName        string  `json:"product_name"`
	ProductDescription string  `json:"product_description"`
	ProductPrice       float64 `json:"product_price"`
	ProductStock       int     `json:"product_stock"`
	ProductDetails     string  `json:"product_details"`
	ProductLink        string  `json:"product_link"`
}
