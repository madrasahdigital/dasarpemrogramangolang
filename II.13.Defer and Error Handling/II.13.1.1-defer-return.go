package main

import "fmt"

func main() {
    pembelianMobil("toyota")
    pembelianMobil("honda")
}

func pembelianMobil(menu string) {
    defer fmt.Println("Terimakasih, silakan tunggu")
    if menu == "toyota" {
        fmt.Print("Pilihan tepat!", " ")
        fmt.Print("toyota solusi berkendaraan!", "\n")
        return
    }

    fmt.Println("Pesanan anda:", menu)
}
