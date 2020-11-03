package main

import "fmt"

func main() {  

    var nilai = 8

    if nilai == 10 {
        fmt.Println("lulus dengan nilai sempurna")
    } else if nilai > 5 {
        fmt.Println("lulus")
    } else if nilai == 4 {
        fmt.Println("hampir lulus")
    } else {
        fmt.Printf("tidak lulus. nilai anda %d\n", nilai)
    }
    
}
