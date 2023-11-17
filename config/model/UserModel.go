package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	FirstName                           string `json:"first_name"`
	LastName                            string `json:"last_name"`
	Email                               string `json:"email" gorm:"unique"`
	MobilePhoneNumber                   string `json:"mobile_phone_number" gorm:"unique"`
	Password                            string `json:"password"`
	RoleID                              int    `json:"role_id" gorm:"foreign_key:RoleID"`
	SubscribedToEmailOffersAndDiscounts bool   `json:"subscribed_to_email_offers_and_discounts"`
	Status                              string `json:"status"`
	TwoFactorAuthentication             bool   `json:"two_factor_authentication"`
	Currency                            int    `json:"currency"`
}
