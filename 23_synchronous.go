package main

import (
	"fmt"
	"time"
)

func main() {
    fmt.Println("Start Synchronous Task...")

    for i := 1; i <= 5; i++ {
        doTask(i)
    }

    fmt.Println("All Synchronous Task Done!")
}

func doTask(i int) {
    fmt.Printf("Start Synchronous Task %d\n", i)
    time.Sleep(1 * time.Second)
    fmt.Printf("End Synchronous Task %d\n", i)
}