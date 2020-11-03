package main

import (
	"fmt"
	"strings"
)


func yourCars(name string, cars ...string) {
    var carsAsString = strings.Join(cars, ", ")

    fmt.Printf("Hello, my name is: %s\n", name)
    fmt.Printf("My cars are: %s\n", carsAsString)
}

func main() {
	yourCars("bruce", "toyota", "honda")
	
	var cars = []string{"toyota", "honda"}
	yourCars("wayne", cars...)
	
}


