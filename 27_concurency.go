package main

import (
	"fmt"
	"time"
)

func main() {
	go count("goroutine 1", 5)
	go count("goroutine 2", 5)

	// Menunggu beberapa waktu agar goroutine memiliki cukup waktu untuk berjalan
	time.Sleep(3 * time.Second)
}

func count(name string, n int) {
	for i := 1; i <= n; i++ {
		fmt.Printf("%s: %d\n", name, i)
		time.Sleep(500 * time.Millisecond)
	}
}