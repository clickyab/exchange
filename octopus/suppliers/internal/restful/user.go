package restful

type User struct {
	IID string `json:"id"`
}

func (u User) ID() string {
	return u.IID
}
