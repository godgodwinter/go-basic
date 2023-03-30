package main

import "fmt"

func fungsigo() {
fnSayHello()
fnParamNamaLengkap("Babeng","Batu")
}

func fnSayHello(){
	fmt.Println("hello world")
}

func fnParamNamaLengkap(firstName string,lastName string)  {
	fmt.Println("Nama Lengkap",firstName,lastName)
}