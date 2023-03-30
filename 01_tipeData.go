package main

import "fmt"

func tipeData() {
// Integer
// 1. int8 range antara -128 - 128 
//  int16 ,int32, int64
// 2. uint8 //unsigned int 0-255
// uint16,uint32,uint64
// 3.Floatin point float32,float64,complex64,complex128
// 4. alias 
// nama lain dari tipe data lainya misal :
// byte nama lain dari uit8
// rune = int32
// int = minimal int32
//  uint = minimal uint32
fmt.Println("Satu",1 )
fmt.Println("angka",3.5 )


// BOOLEAN
// bool true - false
fmt.Println("Benar=",true )

// STRING
// # string adalah kumpulan karakter
// # diawali petik 2 dan diakhiri petik 2 ""
fmt.Println("Ini string")
fmt.Println(len("Ini berapa karakter?")) //20 karakter
fmt.Println("Babeng"[1]) // dihasilkan 97 / byte nya huruf a




}
