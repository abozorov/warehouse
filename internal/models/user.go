package models

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Role     string `json:"role"`
	FullName string `json:"full_name" db:"full_name"`
	Active   bool   `json:"active"`
}

type UserSignIn struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
