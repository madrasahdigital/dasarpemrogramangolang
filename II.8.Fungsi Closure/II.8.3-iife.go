package main

import "fmt"

func main() {
    var angka = []int{6,5,4,8,9,3,2,4,5,7,0,9}

    var newAngka = func(min int) []int {
        var r []int
        for _, e := range angka {
            if e < min {
                continue
            }
            r = append(r, e)
        }
        return r
    }(3)

    fmt.Println("original number :", angka)
    fmt.Println("filtered number :", newAngka)
}
