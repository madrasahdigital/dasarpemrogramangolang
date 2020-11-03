package main

import "fmt"


func main() {  

    var cars = [4]string{"toyota", "mitsubishi", "lexus", "bmw"}

    for i := 0; i < len(cars); i++ {
        fmt.Printf("elemen %d : %s\n", i, cars[i])
    }
    
}
