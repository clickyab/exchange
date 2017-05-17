package dset

import (
	"services/safe"
	"sync"
	"time"
)

var distributedSets = make(map[string]DistributedSet)
var lock = sync.Mutex{}

// GetDistributedSet retrieve DistributedSet from store or make new one if dose not exist
func GetDistributedSet(key string) DistributedSet {
	if distributedSets[key] != nil {
		return distributedSets[key]
	}
	d := &distributedSet{
		key: key,
	}
	return d
}

type distributedSet struct {
	once    sync.Once
	members []string
	adds    []string
	key     string
	exp     time.Time
}

// Members return ads ID
func (d *distributedSet) Members() []string {
	return d.members
}

// Add new ad ID to memebers (after invoking save)
func (d *distributedSet) Add(members ...string) {
	d.adds = append(d.adds, members...)
}

// Key of DistributedSet
func (d *distributedSet) Key() string {
	return d.key
}

// Save added IDs and extend TTL of distributedSet
func (d *distributedSet) Save(t time.Duration) {
	lock.Lock()
	defer lock.Unlock()
	d.exp = time.Now().Add(t)
	d.members = append(d.members, d.adds...)
	distributedSets[d.key] = d
	d.once.Do(func() {
		safe.GoRoutine(d.ttl)
	})
}

func (d *distributedSet) ttl() {
	for {
		<-time.After(time.Until(d.exp))
		if time.Now().Unix() > d.exp.Unix() {
			lock.Lock()
			defer lock.Unlock()
			delete(distributedSets, d.key)
			return
		}
	}
}
