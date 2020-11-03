package main

import "fmt"

type student struct {
    name  string
    grade int
}

func (s student) gantiNama1(name string) {
    fmt.Println("---> on gantiNama1, name changed to", name)
    s.name = name
}

func (s *student) gantiNama2(name string) {
    fmt.Println("---> on gantiNama2, name changed to", name)
    s.name = name
}

func main() {
    var s1 = student{"clark kent", 27}
    fmt.Println("s1 before", s1.name)   

    s1.gantiNama1("bruce wayne")
    fmt.Println("s1 after gantiNama1", s1.name) 

    s1.gantiNama2("peter parker")
    fmt.Println("s1 after gantiNama2", s1.name)  
}
