package main

import (
	"fmt"
	"time")


func main() {
	var ch = make(chan bool)

	time.AfterFunc(4*time.Second, func() {
		fmt.Println("kadaluarsa")
		ch <- true
	})

	fmt.Println("mulai")
	<-ch
	fmt.Println("berakhir")
}
