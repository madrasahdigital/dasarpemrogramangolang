package main

import (
	"fmt"
	"runtime")


func main() {
	runtime.GOMAXPROCS(2)
	var messages = make(chan string)

	go func(who string) {
		var data = fmt.Sprintf("superhero %s", who)
		messages <- data
	}("wayne")

	var message = <-messages
	fmt.Println(message)

}
