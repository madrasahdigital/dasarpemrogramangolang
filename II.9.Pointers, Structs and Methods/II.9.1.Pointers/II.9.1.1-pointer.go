package main

import "fmt"

func findMax(numbers []int, max int) (int, func() []int) {
	var res []int
	for _, e := range numbers {
		if e <= max {
			res = append(res, e)
		}
	}
	return len(res), func() []int {
		return res
	}
}

func main() {
    var angkaA int = 4
    var angkaB *int = &angkaA

    fmt.Println("angkaA (value)   :", angkaA)  // 4
    fmt.Println("angkaA (address) :", &angkaA) // 0xc20800a220

    fmt.Println("angkaB (value)   :", *angkaB) // 4
    fmt.Println("angkaB (address) :", angkaB)  // 0xc20800a220

}
