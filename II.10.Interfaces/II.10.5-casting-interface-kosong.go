package main

import "fmt"

func main() {
	var data map[string]interface{}

	data = map[string]interface{}{
		"name":  "bruce hunt",
		"price": 2000,
		"cars":  []string{"toyota", "manggo", "lexus"},
	}

	fmt.Println(data)
}
