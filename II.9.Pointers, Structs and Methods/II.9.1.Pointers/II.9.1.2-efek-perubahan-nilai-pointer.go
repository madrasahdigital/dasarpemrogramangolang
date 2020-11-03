package main

import "fmt"

func main() {
    var angkaA int = 12
    var angkaB *int = &angkaA

    fmt.Println("angkaA (value)   :", angkaA)
    fmt.Println("angkaA (address) :", &angkaA)
    fmt.Println("angkaB (value)   :", *angkaB)
    fmt.Println("angkaB (address) :", angkaB)

    fmt.Println("")

    angkaA = 9

    fmt.Println("angkaA (value)   :", angkaA)
    fmt.Println("angkaA (address) :", &angkaA)
    fmt.Println("angkaB (value)   :", *angkaB)
    fmt.Println("angkaB (address) :", angkaB)

}
