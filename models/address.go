package models

import "gorm.io/gorm"

type Address struct {
	gorm.Model
	UserID  int    `json:"user_id" gorm:"References:users(ID)"`
	Address string `json:"address"`
}
