package main

import (
	"fmt"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(2)

	var messages = make(chan string)

	var superHero = func(who string) {
		var data = fmt.Sprintf("superhero %s", who)
		messages <- data
	}

	go superHero("bruce wayne")
	go superHero("clark kent")
	go superHero("peter parker")

	var message1 = <-messages
	fmt.Println(message1)

	var message2 = <-messages
	fmt.Println(message2)

	var message3 = <-messages
	fmt.Println(message3)
}
