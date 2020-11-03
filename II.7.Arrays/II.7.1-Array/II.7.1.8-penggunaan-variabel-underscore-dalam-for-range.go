package main

import "fmt"


func main() {  

    var cars = [4]string{"toyota", "mitsubishi", "lexus", "bmw"}

    for _, car := range cars {
        fmt.Printf("elemen : %s\n", car)
    }

}
