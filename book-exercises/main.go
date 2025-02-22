package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
    abort := make(chan bool)
    ticker := time.Tick(time.Second)

    go func() {
        b := []byte{0}
        os.Stdin.Read(b) // blocks
        abort <- true
    }()

    for i := 10; i >= 0; i-- {
        select {
        case <- ticker: fmt.Println(i)
        case <- abort:
            fmt.Println("abort mission!")
            return
        }    
    }
    fmt.Println("launch!")
}