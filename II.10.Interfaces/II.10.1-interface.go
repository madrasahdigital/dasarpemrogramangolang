package main

import (
	"fmt"
	"math"
)


type hitung interface {
	luas() float64
	keliling() float64
}

type hitungLingkaran struct {
	diameter float64
}

func (l hitungLingkaran) jariJari() float64 {
	return l.diameter / 2
}

func (l hitungLingkaran) luas() float64 {
	return math.Pi * math.Pow(l.jariJari(), 2)
}

func (l hitungLingkaran) keliling() float64 {
	return math.Pi * l.diameter
}

type hitungPersegi struct {
	sisi float64
}

func (p hitungPersegi) luas() float64 {
	return math.Pow(p.sisi, 2)
}

func (p hitungPersegi) keliling() float64 {
	return p.sisi * 4
}

func main() {
	var hitungBangunDatar hitung

	hitungBangunDatar = hitungPersegi{10.0}
	fmt.Println("===== hitungPersegi")
	fmt.Println("luas      :", hitungBangunDatar.luas())
	fmt.Println("keliling  :", hitungBangunDatar.keliling())

	hitungBangunDatar = hitungLingkaran{14.0}
	fmt.Println("===== hitungLingkaran")
	fmt.Println("luas      :", hitungBangunDatar.luas())
	fmt.Println("keliling  :", hitungBangunDatar.keliling())
	fmt.Println("jari-jari :", hitungBangunDatar.(hitungLingkaran).jariJari())
}
