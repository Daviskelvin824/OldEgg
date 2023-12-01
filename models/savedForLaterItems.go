package models

import "gorm.io/gorm"

type SavedForLaterItems struct {
	gorm.Model
	UserID    int    `json:"user_id" gorm:"References:users(id)"`
	ProductID string `json:"product_id"`
	Quantity  int    `json:"quantity"`
}
