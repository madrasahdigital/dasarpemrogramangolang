package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

    rand.Seed(time.Now().Unix())
    var angkaAcak int

    angkaAcak = randomRange (2, 10)
    fmt.Println("Angka Acak:", angkaAcak)
    angkaAcak = randomRange (2, 10)
    fmt.Println("Angka Acak:", angkaAcak)
    angkaAcak = randomRange (2, 10)
    fmt.Println("Angka Acak:", angkaAcak)
}

func randomRange(min, max int) int {
    var value = rand.Int() % (max - min + 1) + min
    return value
}
