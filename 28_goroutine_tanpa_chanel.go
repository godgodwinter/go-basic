package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Transaction struct {
	ID       int
	Amount   float64
	Currency string
}

func main() {
	for i := 1; i <= 5; i++ {
		go generateTransaction(i)
	}

	time.Sleep(3 * time.Second)
}

func generateTransaction(id int) {
	transaction := Transaction{
		ID:       id,
		Amount:   rand.Float64() * 1000,
		Currency: "USD",
	}

	processTransaction(transaction)
}

func processTransaction(transaction Transaction) {
	fmt.Printf("Processing transaction ID: %d, Amount: %.2f %s\n", transaction.ID, transaction.Amount, transaction.Currency)
	// Lakukan logika pemrosesan transaksi di sini...
	time.Sleep(1 * time.Second)
}