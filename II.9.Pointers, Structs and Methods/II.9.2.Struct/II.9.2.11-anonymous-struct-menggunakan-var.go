package main

import "fmt"

type person struct {
	name string
	age  int
}

func main() {
	var student struct {
		person
		height int
	}

	student.person = person{"peter", 25}
	student.height = 182

	fmt.Println(student.name)
	fmt.Println(student.age)
	fmt.Println(student.height)
}