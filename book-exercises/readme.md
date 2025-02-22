# tools or exercises from books
the c programming language
the go programming language
the practice of programming


[next](https://learning.oreilly.com/library/view/the-go-programming/9780134190570/ebook_split_084.html)

ideas:
* chat server with features: users entering, user logout, idle user logout
* actor model - build goroutine pipeline with generics, similar to akka
* scheduler, async engine
* context
* go concurence book
* go modules
* go tool

nice lessons:
* `printf` - specify width, nice for formatting output in tables
* simple parsing - `scanf`
* state machines are easy, use them in loops! You can go a long way with them
* obvious, but worth reiterating:
```go
    if aaa && bbb {
    } else {} // == else if !aaa || !bbb <- we can use it to exclude just one condition
```
* image and image/gif packages for raster graphics :D
    * svgs (htmls with formulas) for vector
* use std out in code and redirect to file in shell. Don't need to create files in the code constantly for scripting
    * this is also cool pattern, no need to io.ReadAll and println: `io.Copy(os.stdout, response.Body)`

# concurency

* spinner with go routine, but it's not killed properly. Done channel should be used 
```go
func main() {
    go spinner(100 * time.Millisecond)
    const n = 45
    fibN := fib(n) // slow
    fmt.Printf("\rFibonacci(%d) = %d\n", n, fibN)
}

func spinner(delay time.Duration) {
    for {
        for _, r := range `-\|/` {
            fmt.Printf("\r%c", r)
            time.Sleep(delay)
        }
    }
}
```

* closing channel - sender should close channel!
// You needn’t close every channel when you’ve finished with it. 
// It’s only necessary to close a channel when it is important to 
// tell the receiving goroutines that all data have been sent

* buffered channel - queue, FIFO, send blocks when full, receive blocks when empty
`len` - current state of the queue, `cap` - max allowed elements
* goroutine leaks - when one blocks forever. They should be read till the end. Do not return prematurely from iterated goroutine

```go
errors := make(chan error)
for _, f := range filenames {
    go func(f string) {
        errors <- work(f)
    }(f)
}

for range filenames {
    if err := <-errors; err != nil {
        return err // incorrect: goroutine leak!
    }
}
return nil
```
* select - for multiplexing multiple channels (receiver or sending)
* `time.Tick(1 * time.Second)` leaks goroutine. We need to close it if it's not for the entire lifetime of the app