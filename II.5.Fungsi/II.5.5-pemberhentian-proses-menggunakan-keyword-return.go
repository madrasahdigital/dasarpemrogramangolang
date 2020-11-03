package main

import "fmt"

func main() {
    pembagianAngka(10, 2)
    pembagianAngka(4, 0)
    pembagianAngka(8, -4)
}

func pembagianAngka(m, n int) {
    if n == 0 {
        fmt.Printf("invalid divider. %d cannot divided by %d\n", m, n)
        return
    }

    var res = m / n
    fmt.Printf("%d / %d = %d\n", m, n, res)
}
