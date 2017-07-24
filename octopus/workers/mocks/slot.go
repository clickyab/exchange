package mocks

type Slot struct {
	SWidth, SHeight int
	STRackID        string
	SFallback       string
	Attribute      map[string]string
}

func (s Slot) Attributes() map[string]string {
	return s.Attribute
}

func (s Slot) Width() int {
	return s.SWidth
}

func (s Slot) Height() int {
	return s.SHeight
}

func (s Slot) TrackID() string {
	return s.STRackID
}

func (s Slot) Fallback() string {
	return s.SFallback
}
