package models

type Tutor struct {
	ID             int32      `json:"id"`
	Username       string     `json:"username"`
	FirstName      string     `json:"first_name"`
	LastName       string     `json:"last_name"`
	IsTutor        bool       `json:"is_tutor"`
	Rating         float64    `json:"rating"`
	Subjects       []Subject  `json:"subjects"`
	Email          string     `json:"email"`
	Phone          string     `json:"phone"`
	Contact        string     `json:"contact"`
	About          string     `json:"about"`
	Price          float64    `json:"price"`
	Connections    []*Student `json:"connections"`
	Reviews        []Review   `json:"reviews"`
	ProfilePicture string     `json:"profile_picture"`
	Title          string     `json:"title"`
}
