package restful

type user struct {
	FID string `json:"id"`
}

func (u user) ID() string {
	return u.FID
}
