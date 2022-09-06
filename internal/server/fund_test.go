package server

import (
	"sync"
	"testing"
)

const WORKERS = 10

func BenchmarkWithdrawls(b *testing.B) {
	if b.N < WORKERS {
		return
	}

	amountPerTx := b.N / WORKERS

	server := NewFundServer(b.N)

	var wg sync.WaitGroup

	for i := 0; i < WORKERS; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for i := 0; i < amountPerTx; i++ {
				server.Commands <- WithdrawCommand{Amount: 1}
			}
		}()
	}

	wg.Wait()

	balanceResChannel := make(chan int)
	server.Commands <- BalanceCommand{Response: balanceResChannel}
	balance := <-balanceResChannel

	if balance != 0 {
		b.Error("Expected Balance to be 0, got", balance)
	}
}
