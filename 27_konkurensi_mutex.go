package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	var mutex sync.Mutex

	wg.Add(2)

	go func() {
		defer wg.Done()
		count("goroutine 1", 5, &mutex)
	}()

	go func() {
		defer wg.Done()
		count("goroutine 2", 5, &mutex)
	}()

	wg.Wait()
}

func count(name string, n int, mutex *sync.Mutex) {
	for i := 1; i <= n; i++ {
		mutex.Lock()
		fmt.Printf("%s: %d\n", name, i)
		mutex.Unlock()
		time.Sleep(500 * time.Millisecond)
	}
}