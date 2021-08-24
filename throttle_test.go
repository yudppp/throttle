package throttle_test

import (
	"sync"
	"sync/atomic"
	"testing"
	"time"

	"github.com/yudppp/throttle"
)

const throttleDuration = time.Second * 2
const testInterval = time.Second * 3

func TestThrottle(t *testing.T) {
	throttler := throttle.New(throttleDuration)

	cnt := uint64(0)
	incrementCount := func() {
		atomic.AddUint64(&cnt, 1)
	}

	// once test
	throttler.Do(incrementCount)

	if cnt != 1 {
		t.Errorf("cnt should be 1, but %d", cnt)
	}

	time.Sleep(testInterval)

	// loop test
	for i := 0; i < 10; i++ {
		throttler.Do(incrementCount)
	}

	if cnt != 2 {
		t.Errorf("cnt should be 2, but %d", cnt)
	}

}

func TestThrottleFromGoroutines(t *testing.T) {
	throttler := throttle.New(throttleDuration)
	cnt := uint64(0)
	incrementCount := func() {
		atomic.AddUint64(&cnt, 1)
	}
	var wg sync.WaitGroup
	for i := 0; i < 64; i++ {
		wg.Add(1)
		go func(i int) {
			throttler.Do(incrementCount)
			wg.Done()
		}(i)
	}
	wg.Wait()
	if cnt != 1 {
		t.Errorf("cnt should be 1, but %d", cnt)
	}
}
