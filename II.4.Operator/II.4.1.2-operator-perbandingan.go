package main

import "fmt"

func main() {
	var value = (((4+ 7) % 3) * 8 - 2) / 2
	var isEqual = (value == 2)

	fmt.Printf("nilai %d (%t) \n", value, isEqual)

}
