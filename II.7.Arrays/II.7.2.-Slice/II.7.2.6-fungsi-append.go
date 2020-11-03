package main

import "fmt"

func main() {
	var cars = []string{"toyota", "mitsubishi", "lexus"}
	var cCars = append(cars, "honda")

	fmt.Println(cars)  // ["toyota", "mitsubishi", "lexus"]
	fmt.Println(cCars) // ["toyota", "mitsubishi", "lexus", "honda"]

}
