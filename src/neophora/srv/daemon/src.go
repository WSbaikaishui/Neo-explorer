package daemon

import (
	"sync"
	"time"
)

// T ...
type T struct {
	last time.Time
	lock sync.Mutex
}

// Call ...
func (me *T) Call() {
	me.lock.Lock()
	defer me.lock.Unlock()
	me.last = time.Now()
}

// Task ...
func (me *T) Task() {
	// TODO
}
