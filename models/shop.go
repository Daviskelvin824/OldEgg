package models

import "gorm.io/gorm"

type Shop struct {
	gorm.Model
	ShopName           string `json:"shop_name"`
	ShopEmail          string `json:"shop_email"`
	ShopPassword       string `json:"shop_password"`
	Status             string `json:"status"`
	AboutUs            string `json:"about_us"`
	DisplayPictureLink string `json:"display_picture_link"`
}
