package main

import "fmt"

type Person struct {
    Nama string
    Umur int
}

func main() {
    persons := []Person{
        {Nama: "Sri Kusuma", Umur: 24},
        {Nama: "Budi Susanto", Umur: 28},
        {Nama: "Sri Hartati", Umur: 32},
        {Nama: "Susi Wati", Umur: 27},
        {Nama: "John Doe", Umur: 30},
    }

    filtered := make([]Person, 0)

    for _, person := range persons {
        if startsWith(person.Nama, "Sri") {
            filtered = append(filtered, person)
        }
    }

    fmt.Println(filtered)
}

func startsWith(str, prefix string) bool {
    return len(str) >= len(prefix) && str[0:len(prefix)] == prefix
}