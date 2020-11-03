package library

import "fmt"

func GetAgent(name string) {
	fmt.Println("hello")
	agent(name)
}

func agent(name string) {
	fmt.Println("agent name :", name)
}
