// throttle is function throttling package
package throttle

import (
	"sync"
	"time"
)

// Throttle is an object that will perform exactly one action per duration.
type Throttler interface {
	// Do call the function f if a specified duration has passed
	// since the last function f was called for this instance of Throttle.
	// In other words, given
	// 	var throttle = Throttle.New(time.Minute)
	// if throttle.Do(f) is called multiple times within a minute, only the first call will invoke f,
	// even if f has a different value in each invocation.
	// Waiting for a minute or a new instance of Throttle is required for each function to execute.
	Do(f func())
}

// New is create Throttler instance
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

// Do is Throttler implement
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
