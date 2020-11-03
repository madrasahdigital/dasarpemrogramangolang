package main

import (
	"errors"
	"fmt"
	"strings"
)

func catch() {
    if r := recover(); r != nil {
        fmt.Println("Error occured", r)
    } else {
        fmt.Println("aplikasi berjalan")
    }
}

func validasi(input string) (bool, error) {
    if strings.TrimSpace(input) == "" {
        return false, errors.New("tidak boleh kosong")
    }
    return true, nil
}


func main() {
    defer catch()

    var name string
    fmt.Print("ketik nama anda: ")
    fmt.Scanln(&name)

    if valid, err := validasi(name); valid {
        fmt.Println("halo", name)
    } else {
        panic(err.Error())
        fmt.Println("end")
    }
}
