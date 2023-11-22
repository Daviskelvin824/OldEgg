package models

import "gorm.io/gorm"

type SearchQuery struct {
	gorm.Model
	Email           string `json:"email"`
	Keyword         string `json:"keyword"`
	InnerKeyword    string `json:"inner_keyword"`
	IsAvailableOnly bool   `json:"is_available_only"`
}
