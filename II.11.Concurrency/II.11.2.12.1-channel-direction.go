package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"time"
)

func kirimData(ch chan<- int) {
    for i := 0; true; i++ {
        ch <- i
        time.Sleep(time.Duration(rand.Int()%10+1) * time.Second)
    }
}

func terimaData(ch <-chan int) {
    loop:
    for {
        select {
        case data := <-ch:
            fmt.Print(`data "`, data, `"`, "\n")
        case <-time.After(time.Second * 5):
            fmt.Println("timeout. setelah 5 seconds")
            break loop
        }
    }
}

func main() {
    rand.Seed(time.Now().Unix())
    runtime.GOMAXPROCS(2)

    var messages = make(chan int)

    go kirimData(messages)
    terimaData(messages)
}
