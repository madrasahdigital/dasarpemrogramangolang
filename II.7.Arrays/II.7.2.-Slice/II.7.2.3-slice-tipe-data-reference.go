package main

import "fmt"


func main() {  

    var cars = []string{"toyota", "mitsubishi", "lexus", "bmw"}

	var aCars = cars[0:3]
	var bCars = cars[1:4]

	var aaCars = aCars[1:2]
	var baCars = bCars[0:1]

	fmt.Println(cars)   // [toyota mitsubishi lexus bmw]
	fmt.Println(aCars)  // [toyota mitsubishi lexus]
	fmt.Println(bCars)  // [mitsubishi lexus bmw]
	fmt.Println(aaCars) // [mitsubishi]
	fmt.Println(baCars) // [mitsubishi]

	// Buah "mitsubishi" diubah menjadi "honda"
	baCars[0] = "honda"

	fmt.Println(cars)   // [toyota honda lexus bmw]
	fmt.Println(aCars)  // [toyota honda lexus]
	fmt.Println(bCars)  // [honda lexus bmw]
	fmt.Println(aaCars) // [honda]
	fmt.Println(baCars) // [honda]

} 
