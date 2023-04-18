package models

import "gorm.io/gorm"

type Subject struct {
	gorm.Model `json:"-"`
	Name       string `json:"name" gorm:"unique;column:name"`
}
