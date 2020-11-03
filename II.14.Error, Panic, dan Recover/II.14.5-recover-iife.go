package main

import (
	"fmt"
)

func main() {
    data := []string{"bruce", "peter", "clark"}

    for _, each := range data {

        func() {

             defer func() {
                if r := recover(); r != nil {
                    fmt.Println("Panic penggulangan", each, "| pesan:", r)
                } else {
                    fmt.Println("Aplikasi berjalan")
                }
            }()

            panic("terjadi kesalahan")
        }()

    }
}
