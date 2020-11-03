package main

import "fmt"

func Catch() {
  if r := recover(); r != nil {
    fmt.Println(fmt.Sprintf("Recovered from : %v", r)) 
  }
}

func DoPanic() {
  defer Catch()
  panic("This is panic")
}

func main() {
  DoPanic()
}
