package main

import "fmt"

func map123() {
// map adlaah tipe data kumpulan key dan value
// key nya harus  uniq
// jmlh data bebas
// jika kata kunci sama maka data sebelumnya akan di replase
person := map[string]string{
	"name": "Babeng",
    "age":  "18",
}

fmt.Println(person)
fmt.Println(person["name"])
fmt.Println(person["age"])
person["title"]="Progamer" // untuk menambah data baru di map
fmt.Println(person["title"])

// make(map[tipeData]tipeValue) //untuk membuat map baru
// delete(map,key) menghapus data di map dengan key tertentu

book := make(map[string]string)
book["title"]="Buku Go Lang"
book["harga"]="20000"
fmt.Println(book)
fmt.Println(book["title"])

delete(book, "harga")
fmt.Println(book)
}

