package dset

import (
	"services/safe"
	"sync"
	"time"
)

var distributedSets = make(map[string]DistributedSet)
var lock = sync.Mutex{}

// GetDistributedSet is the distributed set
func GetDistributedSet(key string) DistributedSet {
	if distributedSets[key] != nil {
		return distributedSets[key]
	}
	d := &dist{
		key: key,
	}
	return d
}

type dist struct {
	once    sync.Once
	members []string
	adds    []string
	key     string
	exp     time.Time
}

// Members return
func (d *dist) Members() []string {
	return d.members
}

func (d *dist) Add(members ...string) {
	d.adds = append(d.adds, members...)
}

func (d *dist) Key() string {
	return d.key
}
func (d *dist) Save(t time.Duration) {
	lock.Lock()
	defer lock.Unlock()
	d.exp = time.Now().Add(t)
	d.members = append(d.members, d.adds...)
	distributedSets[d.key] = d
	d.once.Do(func() {
		safe.GoRoutine(d.timer)
	})

}
func (c *dist) timer() {
	for {
		select {
		case <-time.After(time.Until(c.exp)):
			if time.Now().Unix() > c.exp.Unix() {
				lock.Lock()
				defer lock.Unlock()
				delete(distributedSets, c.key)
				return
			}
		}
	}
}
