package main

import (
	"errors"
	"fmt"
	"strings"
)

func validasi(input string) (bool, error) {
    if strings.TrimSpace(input) == "" {
        return false, errors.New("tidak boleh kosong")
    }
    return true, nil
}

func main() {
    var name string
    fmt.Print("ketik nama: ")
    fmt.Scanln(&name)

    if valid, err := validasi(name); valid {
        fmt.Println("hello", name)
    } else {
        fmt.Println(err.Error())
    }
}
