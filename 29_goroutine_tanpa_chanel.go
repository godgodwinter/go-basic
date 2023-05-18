package main

import (
	"fmt"
	"sync"
	"time"
)

type Transaction struct {
	Amount float64
}

func main() {
	saldo := 200000.0

	var wg sync.WaitGroup
	wg.Add(4)

	go processTransaction(10000, &saldo, &wg)
	go processTransaction(20000, &saldo, &wg)
	go processTransaction(180000, &saldo, &wg)
	go processTransaction(40000, &saldo, &wg)

	wg.Wait()

	fmt.Printf("Saldo akhir: %.2f\n", saldo)
}

func processTransaction(amount float64, saldo *float64, wg *sync.WaitGroup) {
	defer wg.Done()

	if amount <= *saldo {
		*saldo -= amount
		fmt.Printf("Transaksi berhasil! Jumlah: %.2f Saldo: %.2f\n", amount, *saldo)
	} else {
		fmt.Printf("Transaksi gagal! Jumlah: %.2f Saldo tidak mencukupi. Saldo: %.2f\n", amount, *saldo)
	}

	time.Sleep(500 * time.Millisecond)
}