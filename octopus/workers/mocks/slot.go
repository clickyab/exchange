package mocks

type Imp struct {
	SWidth, SHeight int
	STRackID        string
	SFallback       string
	Attribute       map[string]string
}

func (s Imp) Attributes() map[string]string {
	return s.Attribute
}

func (s Imp) Width() int {
	return s.SWidth
}

func (s Imp) Height() int {
	return s.SHeight
}

func (s Imp) TrackID() string {
	return s.STRackID
}

func (s Imp) Fallback() string {
	return s.SFallback
}

func (s *Imp) SetAttribute(att string, v string) {
	s.Attribute[att] = v
}
