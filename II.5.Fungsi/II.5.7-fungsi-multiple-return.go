package main

import (
	"fmt"
	"math"
)

func calculate(d float64) (float64, float64) {
    // hitung luas
    var luas = math.Pi * math.Pow(d / 2, 2)
    // hitung keliling
    var lingkaran = math.Pi * d

    // kembalikan 2 nilai
    return luas, lingkaran
}

func main() {
    var diameter float64 = 15
    var luas, lingkaran = calculate(diameter)

    fmt.Printf("luas lingkaran\t\t: %.2f \n", luas)
    fmt.Printf("keliling lingkaran\t: %.2f \n", lingkaran)
}
