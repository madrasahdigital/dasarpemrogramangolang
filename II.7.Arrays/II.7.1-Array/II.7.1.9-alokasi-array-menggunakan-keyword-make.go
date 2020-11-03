package main

import "fmt"


func main() {  

    var cars = make([]string, 2)
    cars[0] = "toyota"
    cars[1] = "mitsubishi"

    fmt.Println(cars)  // [toyota mitsubishi]

}
