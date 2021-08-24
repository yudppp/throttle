# Throttle

![test workflow](https://github.com/yudppp/throttle/actions/workflows/test.yml/badge.svg)
[![Go Report Card](https://goreportcard.com/badge/github.com/yudppp/throttle)](https://goreportcard.com/report/github.com/yudppp/throttle)
[![Coverage Status](https://coveralls.io/repos/github/yudppp/throttle/badge.svg?branch=master)](https://coveralls.io/github/yudppp/throttle?branch=master)

Throttle is an object that will perform exactly one action per duration.
Do call the function f if a specified duration has passed since the last function f was called for this instance of Throttle.

<img width="555" alt="Group 38 (1)" src="https://user-images.githubusercontent.com/4619802/130463248-b27fa321-24ce-47d1-9f9c-5f8b730d73c0.png">

## Examples

### single thread

```go
package main

import (
	"fmt"
	"time"

	"github.com/yudppp/throttle"
)

func main() {
	throttler := throttle.New(time.Second)
	throttler.Do(func() {
		fmt.Println("first call")
	})
	throttler.Do(func() {
		// this function called never.
		fmt.Println("second call")
	})
	time.Sleep(time.Second)
	throttler.Do(func() {
		fmt.Println("third call")
	})
	time.Sleep(time.Second)
}
```

```
$ go run -race main.go
first call
third call
```

### multiple threads

```go
package main

import (
	"fmt"
	"time"

	"github.com/yudppp/throttle"
)

func main() {
	throttler := throttle.New(time.Second)
	var wg sync.WaitGroup
	for i := 0; i < 64; i++ {
		wg.Add(1)
		go func(i int) {
			throttler.Do(func() {
				fmt.Println("called")
			})
			wg.Done()
		}(i)
	}
	wg.Wait()
}
```

```
$ go run -race main.go
called
```

## License

[The MIT License (MIT)](https://github.com/yudppp/throttle/blob/master/LICENSE)
