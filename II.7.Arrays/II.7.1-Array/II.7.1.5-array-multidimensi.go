package main

import "fmt"


func main() {  
    var angka1 = [2][3]int{[3]int{3, 2, 3}, [3]int{3, 4, 5}}
    var angka2 = [2][3]int{{3, 2, 3}, {3, 4, 5}}

    fmt.Println("angka1", angka1)
    fmt.Println("angka2", angka2)

}
