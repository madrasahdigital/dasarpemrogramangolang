package main

import "fmt"

func main() {
    var angka = 7
    fmt.Println("before :", angka) // 7

    change(&angka, 9)
    fmt.Println("after  :", angka) // 9
}

func change(original *int, value int) {
    *original = value
}
