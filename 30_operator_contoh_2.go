package main

import "fmt"

func updateValue(ptr *int) {
    *ptr = 20
}

func main() {
    x := 10
    fmt.Println("Nilai x sebelum diupdate:", x)
    
    updateValue(&x)
    
    fmt.Println("Nilai x setelah diupdate:", x)
}