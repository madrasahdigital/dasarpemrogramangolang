package main

import "fmt"

func main() {  

    var harga = 8840.0

    if persen := harga / 100; persen >= 100 {
        fmt.Printf("%.1f%s perfect!\n", persen, "%")
    } else if persen >= 70 {
        fmt.Printf("%.1f%s good\n", persen, "%")
    } else {
        fmt.Printf("%.1f%s not bad\n", persen, "%")
    }
    
}
