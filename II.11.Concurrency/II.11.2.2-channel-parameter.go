package main

import (
	"fmt"
	"runtime"
)

func cetakPesan(what chan string) {
    fmt.Println(<-what)
}  

func main() {
    runtime.GOMAXPROCS(2)

    var messages = make(chan string)

    for _, each := range []string{"batman", "superman", "spiderman"} {
        go func(who string) {
            var data = fmt.Sprintf("superhero %s", who)
            messages <- data
        }(each)
    }

    for i := 0; i < 3; i++ {
        cetakPesan(messages)
    }
}
