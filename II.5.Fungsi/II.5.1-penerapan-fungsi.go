package main

import (
	"fmt"
	"strings")


func main() {
    var names = []string{"Bruce", "Wayne"}
    cetakPesan("hello", names)
}

func cetakPesan(message string, arr []string) {
    var nameString = strings.Join(arr, " ")
    fmt.Println(message, nameString)
}
