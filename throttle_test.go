package throttle_test

import (
	"testing"
	"time"

	"github.com/yudppp/throttle"
)

func TestThrottle(t *testing.T) {
	throttler := throttle.New(time.Second)

	cnt := 0
	// once test
	throttler.Do(func() {
		cnt++
	})

	if cnt != 1 {
		t.Errorf("cnt should be 1, but %d", cnt)
	}

	time.Sleep(time.Second)

	// loop test
	for i := 0; i < 10; i++ {
		throttler.Do(func() {
			cnt++
		})
	}

	if cnt != 2 {
		t.Errorf("cnt should be 2, but %d", cnt)
	}

	time.Sleep(time.Second)

	// conflict test
	for i := 0; i < 10; i++ {
		go func() {
			for j := 0; j < 10; j++ {
				throttler.Do(func() {
					cnt++
				})
			}
		}()
	}

	if cnt != 3 {
		t.Errorf("cnt should be 3, but %d", cnt)
	}
}
