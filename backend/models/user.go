package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username       string    `json:"username"`
	Password       string    `json:"password"`
	FirstName      string    `json:"first_name"`
	LastName       string    `json:"last_name"`
	IsTutor        bool      `json:"is_tutor"`
	Rating         float64   `json:"rating"`
	Subjects       []Subject `gorm:"many2many:user_subjects" json:"subjects"`
	Email          string    `json:"email"`
	Phone          string    `json:"phone"`
	Contact        string    `json:"contact"`
	About          string    `json:"about"`
	Grade          int32     `json:"grade"`
	Connections    []*User   `gorm:"many2many:user_connections" json:"connections"`
	Price          float64   `json:"price"`
	Reviews        []Review  `gorm:"many2many:user_reviews" json:"reviews"`
	ProfilePicture string    `json:"profile_picture"`
	Title          string    `json:"title"`
}
