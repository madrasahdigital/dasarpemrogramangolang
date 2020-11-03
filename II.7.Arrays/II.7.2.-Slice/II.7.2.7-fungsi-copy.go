package main

import "fmt"


func main() {  

    dst := make([]string, 3)
    src := []string{"kia", "honda", "toyota", "mitsubishi"}
    n := copy(dst, src)

    fmt.Println(dst) // kia honda toyota
    fmt.Println(src) // kia honda toyota mitsubishi
    fmt.Println(n)   // 3

	dst := []string{"toyota", "toyota", "toyota"}
    src := []string{"kia", "honda"}
    n := copy(dst, src)
    fmt.Println(dst) // kia honda toyota
    fmt.Println(src) // kia honda
    fmt.Println(n)   // 2


}