package user

import "time"

type activeStatus string

const (
	activeStatusTrue  activeStatus = "yes"
	activeStatusFalse activeStatus = "no"
)

// User user model in database
// @Model {
//		table = users
//		primary = true, id
//		find_by = id
//		transaction = insert
//		list = yes
// }
type User struct {
	ID        int          `json:"id" db:"id"`
	Email     string       `json:"email" db:"email"`
	Domain    string       `json:"domain" db:"domain"`
	Password  string       `json:"password" db:"password"`
	Active    activeStatus `json:"active" db:"active"`
	CreatedAt *time.Time   `json:"created_at"  db:"created_at"`
	UpdatedAt *time.Time   `json:"updated_at" db:"updated_at"`
}
