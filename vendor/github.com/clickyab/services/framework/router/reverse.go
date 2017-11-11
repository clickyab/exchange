package router

import (
	"sync"

	"fmt"

	"strings"

	"github.com/clickyab/services/assert"
	"github.com/sirupsen/logrus"
)

var (
	reverse = map[string]string{}
	lock    = sync.RWMutex{}
)

func addRoute(name, path string) {
	lock.Lock()
	defer lock.Unlock()

	if p, ok := reverse[name]; ok {
		logrus.Panicf("duplicate name %s, already registered for %s and requested for %s", name, p, path)
	}

	reverse[name] = path
}

// Path return the path for this route if its already registered
func Path(name string, params map[string]string, catch ...string) (string, error) {
	lock.RLock()
	defer lock.RUnlock()

	p, ok := reverse[name]
	if !ok {
		return "", fmt.Errorf("no route with name %s", name)
	}

	parts := strings.Split(p, "/")
	res := []string{}
	for i := range parts {
		if parts[i] == "" {
			continue
		}
		assert.True(len(parts[i]) > 0, parts[i])
		mm := parts[i]
		if mm[0] == ':' {
			d, ok := params[mm[1:]]
			if !ok {
				return "", fmt.Errorf("can not find parameter %s in all data", mm[1:])
			}
			mm = d
		} else if mm[0] == '*' {
			// catch all parameter
			res = append(res, catch...)
			break
		}
		res = append(res, mm)
	}

	return "/" + strings.Join(res, "/"), nil

}
