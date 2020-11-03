package main

import "fmt"


func main() {  
    var names [4]string
    names[0] = "paris"
    names[1] = "van"
    names[2] = "java"
    names[3] = "bandung"
    names[4] = "jawa barat" // baris kode ini menghasilkan error

    fmt.Println(names[0], names[1], names[2], names[3])
    
}
