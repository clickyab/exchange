package openrtb

type user struct {
	IID string `json:"id"`
}

func (u user) ID() string {
	return u.IID
}

//TODO remove just for lint
func init() {
	if false {
		panic(user{})
	}
}
