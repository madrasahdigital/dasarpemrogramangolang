package main

import "fmt"


func main() {  

    var cars = [4]string{"toyota", "mitsubishi", "lexus", "bmw"}

    for i, car := range cars {
        fmt.Printf("elemen %d : %s\n", i, car)
    }

}
