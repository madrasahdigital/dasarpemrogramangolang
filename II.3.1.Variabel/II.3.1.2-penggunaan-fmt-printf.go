package main

import "fmt"

func main() {

	var firstName string = "bruce"

	var lastName string
	lastName = "wayne"

	fmt.Printf("halo bruce wayne!\n")
	fmt.Printf("halo %s %s!\n", firstName, lastName)
	fmt.Println("halo", firstName, lastName+"!")
	
}
