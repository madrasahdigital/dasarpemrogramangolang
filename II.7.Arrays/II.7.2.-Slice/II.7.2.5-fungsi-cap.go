package main

import "fmt"


func main() {  
    
    var cars = []string{"toyota", "mitsubishi", "lexus", "bmw"}	
    fmt.Println(len(cars))  // len: 4
    fmt.Println(cap(cars))  // cap: 4

    var aCars = cars[0:3]
    fmt.Println(len(aCars)) // len: 3
    fmt.Println(cap(aCars)) // cap: 4

    var bCars = cars[1:4]
    fmt.Println(len(bCars)) // len: 3
    fmt.Println(cap(bCars)) // cap: 3

}
