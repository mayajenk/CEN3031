package models

import "gorm.io/gorm"

type Review struct {
	gorm.Model
	ReviewerID uint    `gorm:"index" json:"reviewer_id"`
	RevieweeID uint    `gorm:"index" json:"reviewee_id"`
	ReviewText string  `json:"review_text"`
	Rating     float64 `json:"rating"`
}
