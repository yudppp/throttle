package throttle_test

import (
	"testing"
	"time"

	"github.com/yudppp/throttle"
)

const throttleDuration = time.Second * 5
const testInterval = time.Second * 6

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

	time.Sleep(testInterval)

	// loop test
	for i := 0; i < 10; i++ {
		throttler.Do(func() {
			cnt++
		})
	}

	if cnt != 2 {
		t.Errorf("cnt should be 2, but %d", cnt)
	}

	time.Sleep(testInterval)

	// conflict test
	for i := 0; i < 8; i++ {
		go func() {
			for j := 0; j < 8; j++ {
				throttler.Do(func() {
					cnt++
				})
			}
		}()
	}

	time.Sleep(testInterval)

	if cnt != 3 {
		t.Errorf("cnt should be 3, but %d", cnt)
	}
}
