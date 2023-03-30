package main

import "fmt"

func constant() {
const firstName  ="Babeng"
// firstName="bobo" //errorr

// constant tidak harus di gunakan 
// fmt.Println(firstName)

const (
	name="ini nama"
	age=20
)

fmt.Println(name,age,firstName)
}
