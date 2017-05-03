package config

import (
	"sync"
	"time"

	"github.com/Sirupsen/logrus"
	"gopkg.in/fzerorubigd/onion.v2"
)

type variable struct {
	ref, def interface{}
	key      string
}

var (
	allVariables []variable
	lock         = &sync.Mutex{}
)

// RegisterString add an string to config
func RegisterString(key string, def string, description string) *string {
	lock.Lock()
	defer lock.Unlock()
	var ref string

	setDescription(key, description)
	allVariables = append(allVariables, variable{
		ref: &ref,
		def: def,
		key: key,
	})

	return &ref
}

// RegisterInt add an int to config
func RegisterInt(key string, def int, description string) *int {
	lock.Lock()
	defer lock.Unlock()
	var ref int

	setDescription(key, description)
	allVariables = append(allVariables, variable{
		ref: &ref,
		def: def,
		key: key,
	})

	return &ref
}

// RegisterInt64 add an int to config
func RegisterInt64(key string, def int64, description string) *int64 {
	lock.Lock()
	defer lock.Unlock()
	var ref int64
	setDescription(key, description)

	allVariables = append(allVariables, variable{
		ref: &ref,
		def: def,
		key: key,
	})

	return &ref
}

// RegisterFloat64 add an int to config
func RegisterFloat64(key string, def float64, description string) *float64 {
	lock.Lock()
	defer lock.Unlock()
	var ref float64
	setDescription(key, description)

	allVariables = append(allVariables, variable{
		ref: &ref,
		def: def,
		key: key,
	})

	return &ref
}

// RegisterBoolean add a boolean to config
func RegisterBoolean(key string, def bool, description string) *bool {
	lock.Lock()
	defer lock.Unlock()
	var ref bool
	setDescription(key, description)

	allVariables = append(allVariables, variable{
		ref: &ref,
		def: def,
		key: key,
	})

	return &ref
}

// RegisterDuration add an duration to config
func RegisterDuration(key string, def time.Duration, description string) *time.Duration {
	lock.Lock()
	defer lock.Unlock()
	var ref time.Duration
	setDescription(key, description)

	allVariables = append(allVariables, variable{
		ref: &ref,
		def: def,
		key: key,
	})

	return &ref
}

func load(o *onion.Onion) {
	lock.Lock()
	defer lock.Unlock()

	for i := range allVariables {
		switch t := allVariables[i].def.(type) {
		case string:
			v := o.GetStringDefault(allVariables[i].key, t)
			vs := allVariables[i].ref.(*string)
			*vs = v
		case bool:
			v := o.GetBoolDefault(allVariables[i].key, t)
			vs := allVariables[i].ref.(*bool)
			*vs = v
		case int:
			v := o.GetIntDefault(allVariables[i].key, t)
			vs := allVariables[i].ref.(*int)
			*vs = v
		case int64:
			v := o.GetInt64Default(allVariables[i].key, t)
			vs := allVariables[i].ref.(*int64)
			*vs = v
		case float64:
			v := o.GetFloat64Default(allVariables[i].key, t)
			vs := allVariables[i].ref.(*float64)
			*vs = v
		case time.Duration:
			v := o.GetDurationDefault(allVariables[i].key, t)
			vs := allVariables[i].ref.(*time.Duration)
			*vs = v
		default:
			logrus.Panic("wtf :/")
		}
	}
}
