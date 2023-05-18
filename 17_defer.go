package main

import (
	"fmt"
)

// func endApp(){
// 	fmt.Println("fn End App")
// }

// func runApp(error bool){
// 	defer endApp()
// 	if error{
// 		panic("ERRROR")
// 	}
// }

func main() {
	fmt.Println("Mulai")
	defer fmt.Println("Pernyataan ini ditunda1")
	defer fmt.Println("Pernyataan ini ditunda2")
	// runApp(true)
	fmt.Println("Selesai")
}