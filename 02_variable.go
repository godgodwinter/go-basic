package main

import "fmt"

func variable() {
// VARIABLE
// tmpat untuk menyimpan data 
fmt.Println("VARIABLE")
var name string
name = "Babeng ABC"
fmt.Println(name)
name = "Babeng saja"
fmt.Println(name)
// name=123 //error karena string dimasukan integer

var alamat="Malang" //otomatis string 
fmt.Println(alamat)

var umur int8=30
fmt.Println(umur)

var tinggi =165
fmt.Println(tinggi)
// var tidak wajib 
// alternatif deklarasinya bisa gunakan :=
// ex:
berat := 65 //alternatif untuk deklarasi variable
fmt.Println(berat)
berat =64 

// alternatif deklarasi variable 
var (
	firstName="Paijo"
	lastName="Pikachu"
)
fmt.Println(firstName)
fmt.Println(lastName)


}
