![test workflow](https://github.com/yudppp/throttle/actions/workflows/test.yml/badge.svg)
# throttle

Throttle is an object that will perform exactly one action per duration.

```go
var throttler = throttle.New(time.Second*5)

func SomeFunc() {
    throttler.Do(func(){
        fmt.Println("run")
    })
}
```

## License

[The MIT License (MIT)](https://github.com/yudppp/throttle/blob/master/LICENSE)
