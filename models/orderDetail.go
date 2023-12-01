package models

import "gorm.io/gorm"

type OrderDetail struct {
	gorm.Model
	OrderHeaderID uint   `json:"order_header_id" gorm:"References:order_headers(ID)"`
	ProductID     string `json:"product_id" gorm:"References:products(product_id)"`
	Quantity      int    `json:"quantity"`
	Status        string `json:"status"`
}
