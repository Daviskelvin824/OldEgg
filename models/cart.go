package models

import "gorm.io/gorm"

type Cart struct {
	gorm.Model
	UserID    int    `json:"user_id" gorm:"References:users(ID)"`
	ProductID string `json:"product_id"`
	Quantity  int    `json:"quantity"`
}
