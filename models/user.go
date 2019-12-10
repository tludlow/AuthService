package models

//User - A user entity from the database table 'users'
type User struct {
	Username  string `json:"username"`
	Email     string `json:"email"`
	CreatedAt string `json:"created_at"`
}
