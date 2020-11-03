package main

import "fmt"

type person struct {
	name string
	age  int
}

func main() {
	var allStudents = []person{
		{name: "Bruce", age: 23},
		{name: "Clark", age: 23},
		{name: "Peter", age: 22},
	}
	
	for _, student := range allStudents {
		fmt.Println(student.name, "age is", student.age)
	}
}
