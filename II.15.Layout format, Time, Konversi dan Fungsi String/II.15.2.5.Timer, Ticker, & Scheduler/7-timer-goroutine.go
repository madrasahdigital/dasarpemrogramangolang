package main

import (
	"fmt"
	"os"
	"time"
)


func timer(timeout int, ch chan<- bool) {
	time.AfterFunc(time.Duration(timeout)*time.Second, func() {
		ch <- true
	})
}

func watcher(timeout int, ch <-chan bool) {
	<-ch
	fmt.Println("\ntime out! tidak ada respon", timeout, "seconds")
	os.Exit(0)
}

func main() {
	var timeout = 5
	var ch = make(chan bool)

	go timer(timeout, ch)
	go watcher(timeout, ch)

	var input string
	fmt.Print("berapa 725/25 ? ")
	fmt.Scan(&input)

	if input == "29" {
		fmt.Println("benar!")
	} else {
		fmt.Println("salah!")
	}
}
