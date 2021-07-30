# throttle

throttling function called
sync.Once like interface

```go
var throttler = throttle.New(time.Second*5)

func SomeFunc() {
    throttler.Do(func(){
        fmt.Println("run")
    })
}
```


## License

[The MIT License (MIT)](http://yudppp.mit-license.org/)
