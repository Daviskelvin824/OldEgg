package models

import "gorm.io/gorm"

type PaymentMethod struct {
	gorm.Model
	PaymentMethodName string `json:"payment_method_name"`
}
