package main

import (
	"fmt"
)

func main() {
	number := 3

	if number == 3 {
		fmt.Println("hello 1")
		func() {
			defer fmt.Println("hello 3")
		}()
	}

	fmt.Println("hello 2")
}
