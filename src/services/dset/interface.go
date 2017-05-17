package dset

import "time"

// DistributedSet is a set distributed
type DistributedSet interface {
	// Members return all members of this set
	Members() []string
	// Add new item to set
	Add(...string)
	// Key return the master key
	Key() string
	// Save the set with lifetime
	Save(time.Duration)
}

// GetDistributedSet retrieve DistributedSet from store or make new one if dose not exist
func GetDistributedSet(key string) DistributedSet {
	return nil
}
