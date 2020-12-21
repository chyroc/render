package render

import (
	"reflect"
	"sync"
)

var formatMap map[reflect.Type]func(v interface{}) string
var lock sync.RWMutex

func init() {
	formatMap = map[reflect.Type]func(v interface{}) string{}
}

func RegisterFormat(val interface{}, format func(v interface{}) string) {
	if format == nil {
		return
	}

	lock.Lock()
	defer lock.Unlock()

	v := reflect.ValueOf(val)
	if !v.IsValid() {
		return
	}
	formatMap[v.Type()] = format
}

func getFormatFunc(val reflect.Type) func(v interface{}) string {
	lock.RLock()
	defer lock.RUnlock()

	return formatMap[val]
}
