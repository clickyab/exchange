package aaa

// User for user table in db
type User struct {
	Username string `json:"username" db:"username"`
	Password string `json:"password" db:"password"`
}
