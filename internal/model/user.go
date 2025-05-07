package model

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Role     string `json:"role"` //basic, admin
}
