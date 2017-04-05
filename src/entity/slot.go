package entity

// Slot is the slot of the app
type Slot interface {
	// Size return the primary size of this slot
	Width() int
	Height() int
	// StateID is an string for this slot, its a random at first but the value is not changed at all other calls
	StateID() string
}
