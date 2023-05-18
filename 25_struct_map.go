package main

import (
	"fmt"
)


func main() {
	// Contoh Struct
	type Person struct {
		ID   int
		Name string
		Age  int
	}
	// Inisialisasi struct
	person := Person{
		ID:   1,
		Name: "John Doe",
		Age:  30,
	}

	fmt.Println(person)

	// Map dengan Struct
	people := map[int]Person{
		1: {ID: 1, Name: "John Doe", Age: 30},
		2: {ID: 2, Name: "Jane Smith", Age: 25},
		3: {ID: 3, Name: "Mike Johnson", Age: 35},
	}

	// Mengakses data dari map
	fmt.Println(people[1])

	// Mengisi dan Mengupdate data where ID
	people[4] = Person{ID: 4, Name: "Sarah Adams", Age: 28}
	people[4] = Person{ID: 4, Name: "Sarah Anderson", Age: 30}

	fmt.Println(people[4])

	// Menghapus data where ID
	delete(people, 4)
	fmt.Println(people)

	// Operasi pop dan slice
	slice := []int{1, 2, 3, 4, 5}

	// Pop elemen terakhir
	popped := slice[len(slice)-1]
	slice = slice[:len(slice)-1]
	fmt.Println("Popped:", popped)
	fmt.Println("Slice:", slice)

	// Membuat slice baru dari slice awal
	newSlice := slice[1:3]
	fmt.Println("New Slice:", newSlice)
}