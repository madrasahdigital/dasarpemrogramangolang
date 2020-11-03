package main

import "fmt"

func main() {
	var kiri = false
	var kanan = true

	var kiriDanKanan = kiri && kanan
	fmt.Printf("kiri && kanan \t(%t) \n", kiriDanKanan)

	var kiriOrkanan = kiri || kanan
	fmt.Printf("kiri || kanan \t(%t) \n", kiriOrkanan)

	var kiriReverse = !kiri
	fmt.Printf("!kiri \t\t(%t) \n", kiriReverse)

}
