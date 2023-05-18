package main

import "fmt"

func executeFunction(f func()) {
	f()
}

func main() {
	executeFunction(func() {
		fmt.Println("Hello, anonymous function!")
	})
}