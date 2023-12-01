package models

import "gorm.io/gorm"

type OrderHeader struct {
	gorm.Model
	AddressID          uint `json:"address_id" gorm:"References:addresses(ID)"`
	DeliveryProviderID uint `json:"delivery_provider_id" gorm:"References:delivery_providers(ID)"`
	PaymentMethodID    uint `json:"payment_method_id" gorm:"References:payment_methods(ID)"`
	UserID             uint `json:"user_id" gorm:"References:users(ID)"`
}
