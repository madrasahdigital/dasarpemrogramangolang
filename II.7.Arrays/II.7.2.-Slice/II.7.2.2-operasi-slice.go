package main

import "fmt"


func main() {  

    var cars = []string{"toyota", "mitsubishi", "lexus", "bmw"}
	var newCars = cars[0:2]
	fmt.Println(newCars) // ["toyota", "mitsubishi"]

} 
