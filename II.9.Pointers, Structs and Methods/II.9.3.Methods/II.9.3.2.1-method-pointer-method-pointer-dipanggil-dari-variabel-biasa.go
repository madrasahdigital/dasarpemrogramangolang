package main

import "fmt"

type student struct {
	name  string
	grade int
}

func (s *student) sayHello() {
	fmt.Println("hello", s.name)
}

func main() {
	
	var s1 = student{"clark kent", 27}
	s1.sayHello()

	var s2 = &student{"bruce wayne", 28}
	s2.sayHello()

}
