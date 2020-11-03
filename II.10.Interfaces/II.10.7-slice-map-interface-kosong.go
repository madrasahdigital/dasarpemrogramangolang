package main

import "fmt"

func main() {
	var person = []map[string]interface{}{
		{"name": "Wayne", "age": 23},
		{"name": "Clark", "age": 23},
		{"name": "Peter", "age": 22},
	}
	
	for _, each := range person {
		fmt.Println(each["name"], "age is", each["age"])
	}
	
	var cars = []interface{}{
		map[string]interface{}{"name": "toyota", "total": 10},
		[]string{"honda", "mitsubishi", "honda"},
		"lexus",
	}
	
	for _, each := range cars {
		fmt.Println(each)
	}

}
