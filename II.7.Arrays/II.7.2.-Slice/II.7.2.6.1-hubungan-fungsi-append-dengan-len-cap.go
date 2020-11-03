package main

import "fmt"

func main() {
	var cars = []string{"toyota", "mitsubishi", "lexus"}
	var bCars = cars[0:2]

	fmt.Println(cap(bCars)) // 3
	fmt.Println(len(bCars)) // 2

	fmt.Println(cars)  // ["toyota", "mitsubishi", "lexus"]
	fmt.Println(bCars) // ["toyota", "mitsubishi"]

	var cCars = append(bCars, "honda")

	fmt.Println(cars)  // ["toyota", "mitsubishi", "honda"]
	fmt.Println(bCars) // ["toyota", "mitsubishi"]
	fmt.Println(cCars) // ["toyota", "mitsubishi", "honda"]


}
