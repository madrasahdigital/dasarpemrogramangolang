package main

import "fmt"
func main() {  

    var cars = []string{"toyota", "mitsubishi", "lexus"}
    var aCars = cars[0:2]
    var bCars = cars[0:2:2]

    fmt.Println(cars)      // ["toyota", "mitsubishi", "lexus"]
    fmt.Println(len(cars)) // len: 3
    fmt.Println(cap(cars)) // cap: 3

    fmt.Println(aCars)      // ["toyota", "mitsubishi"]
    fmt.Println(len(aCars)) // len: 2
    fmt.Println(cap(aCars)) // cap: 3

    fmt.Println(bCars)      // ["toyota", "mitsubishi"]
    fmt.Println(len(bCars)) // len: 2
    fmt.Println(cap(bCars)) // cap: 2

}
