package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
    fmt.Println("Start Asynchronous Task...")

    var wg sync.WaitGroup
    for i := 1; i <= 5; i++ {
        wg.Add(1)
        go doTaskAsync(i, &wg)
    }

    wg.Wait()

    fmt.Println("All Asynchronous Task Done!")
}

func doTaskAsync(i int, wg *sync.WaitGroup) {
    fmt.Printf("Start Asynchronous Task %d\n", i)
    time.Sleep(1 * time.Second)
    fmt.Printf("End Asynchronous Task %d\n", i)
    wg.Done()
}