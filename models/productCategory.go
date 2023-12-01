package models

import "gorm.io/gorm"

type ProductCategory struct {
	gorm.Model
	ProductCategoryID   int    `json:"product_category_id" gorm:"primary_key"`
	ProductCategoryName string `json:"product_category_name"`
}
