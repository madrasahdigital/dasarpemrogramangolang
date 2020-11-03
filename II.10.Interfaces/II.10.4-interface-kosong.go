package main

import "fmt"

func main() {
	var hero interface{}

	hero = "clark kent"
	fmt.Println(hero)

	hero = []string{"superman", "batman", "spiderman"}
	fmt.Println(hero)

	hero = 12.4
	fmt.Println(hero)
}
