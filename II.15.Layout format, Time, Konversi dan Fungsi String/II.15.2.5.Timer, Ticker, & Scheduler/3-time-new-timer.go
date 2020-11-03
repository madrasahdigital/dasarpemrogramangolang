package main

import (
	"fmt"
	"time")


func main() {
	var timer = time.NewTimer(4 * time.Second)
	fmt.Println("mulai")
	<-timer.C
	fmt.Println("selesai")
}
