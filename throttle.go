package throttle

import (
	"sync"
	"time"
)

// Throttler is throttling function call
type Throttler interface {
	Do(f func())
}

// New is New Throttler Instance
func New(duration time.Duration) Throttler {
	return &throttle{
		duration: duration,
	}
}

type throttle struct {
	duration time.Duration
	once     sync.Once
	m        sync.Mutex
}

// Do is throttling function call
func (t *throttle) Do(f func()) {
	t.m.Lock()
	defer t.m.Unlock()
	t.once.Do(func() {
		go func() {
			time.Sleep(t.duration)
			t.m.Lock()
			defer t.m.Unlock()
			t.once = sync.Once{}
		}()
		f()
	})
}
