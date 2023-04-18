package models

type Student struct {
	ID             int32    `json:"id"`
	Username       string   `json:"username"`
	FirstName      string   `json:"first_name"`
	LastName       string   `json:"last_name"`
	IsTutor        bool     `json:"is_tutor"`
	Rating         float64  `json:"rating"`
	Email          string   `json:"email"`
	Phone          string   `json:"phone"`
	Contact        string   `json:"contact"`
	About          string   `json:"about"`
	Grade          int32    `json:"grade"`
	Connections    []*Tutor `json:"connections"`
	Reviews        []Review `json:"reviews"`
	ProfilePicture string   `json:"profile_picture"`
}
