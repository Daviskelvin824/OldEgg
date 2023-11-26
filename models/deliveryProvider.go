package models

import "gorm.io/gorm"

type DeliveryProvider struct {
	gorm.Model
	DeliveryProviderName string `json:"delivery_provider_name"`
}
