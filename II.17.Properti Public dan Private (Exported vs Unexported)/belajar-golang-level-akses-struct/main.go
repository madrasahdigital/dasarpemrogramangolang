package main

import (
	"fmt"

	"II.17.2.belajar-golang-level-akses-struct/library"
)

func main() {
	var s1 = library.Student{"bruce", 25}
	fmt.Println("name ", s1.Name)
	fmt.Println("grade", s1.Grade)
}
