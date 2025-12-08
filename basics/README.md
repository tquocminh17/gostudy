NOTE:
1. Slice does not contain the value, but the address of the initial array/slice
2. Map always refer to address of the underlying map
3. variables / struct is copied by value, not reference address, except using pointers
4. return value, error instead of throwing panic by return fmt.Errorf("Failed")
5. method of the struct
6. Use empty interface (var j interface{} = [3]int{1, 2, 3}) on the fly, before using it, you should figure out what object is it before doing sth useful
7. BEST PRACTICE:
    - Use many, small interfaces
    - Single method interfaces are some of the most powerful and flexible
        - io.Writer, io.Reader, interface{}
    - Don't export interfaces for types that will be consumed
    - Do export interfaces for types that will be used by package
    - Design functions and methods to receive interfaces whenever possible

---
GO ROUTINES: abstraction of threads
1. Don't create goroutines in libraries. Let's consumer control concurrency
2. When creating a goroutine, know how it will end. Avoids subtle memory leaks
3. Check for race conditions at compile time (go run -race src/main.go)
4. Too many threads can slow down the application, let's benchmark with different threads options

---
CHANNEL is design for data transmission between goroutines

```
// go mod init <module-name>
go mod download
// go run <filename>.go

go build
./gostudy
```
