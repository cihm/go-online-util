package concurrent

import (
	"sync"
)

type SafeMap struct {
	sync.RWMutex
	Map map[string]int
}

func NewSafeMap() *SafeMap {
	sm := new(SafeMap)
	sm.Map = make(map[string]int)
	return sm

}

func (sm *SafeMap) ReadMap(key string) int {
	sm.RLock()
	value := sm.Map[key]
	sm.RUnlock()
	return value
}

func (sm *SafeMap) WriteMap(key string, value int) {
	sm.Lock()
	sm.Map[key] = value
	sm.Unlock()
}
