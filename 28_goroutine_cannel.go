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
	transactions := make(chan Transaction)

	go generateTransactions(transactions)
	go processTransactions(transactions)

	time.Sleep(3 * time.Second)
}

func generateTransactions(transactions chan<- Transaction) {
	for i := 1; i <= 5; i++ {
		transaction := Transaction{
			ID:       i,
			Amount:   rand.Float64() * 1000,
			Currency: "USD",
		}
		transactions <- transaction
		time.Sleep(500 * time.Millisecond)
	}
	close(transactions)
}

func processTransactions(transactions <-chan Transaction) {
	for transaction := range transactions {
		fmt.Printf("Processing transaction ID: %d, Amount: %.2f %s\n", transaction.ID, transaction.Amount, transaction.Currency)
		// Lakukan logika pemrosesan transaksi di sini...
		time.Sleep(1 * time.Second)
	}
}