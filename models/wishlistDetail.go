package models

import "gorm.io/gorm"

type WishlistDetail struct {
	gorm.Model
	WishlistDetailID int64  `json:"id" gorm:"References:wishlists(ID)"`
	ProductID        string `json:"product_id"`
	Quantity         int    `json:"quantity"`
}
