package main

import (
	"fmt"
	"os"
)


func main() {
    defer fmt.Println("hello")
    os.Exit(1)
    fmt.Println("welcome")
}
