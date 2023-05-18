package main

import "fmt"

// Contoh interface
type Animal interface {
	Sound() string
}

// Implementasi interface
type Dog struct{}

func (d Dog) Sound() string {
	return "Woof!"
}

type Cat struct{}

func (c Cat) Sound() string {
	return "Meow!"
}

func main() {
	// Membuat map dengan interface
	animals := make(map[string]Animal)
	animals["dog"] = Dog{}
	animals["cat"] = Cat{}

	// Mengakses dan menggunakan objek dalam map menggunakan interface
	for key, value := range animals {
		fmt.Printf("%s: %s\n", key, value.Sound())
	}
}