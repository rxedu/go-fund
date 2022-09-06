package server

import (
	"sync"
	"testing"

	"github.com/rxedu/go-fund/internal/model"
)

const WORKERS = 10

func BenchmarkWithdrawls(b *testing.B) {
	if b.N < WORKERS {
		return
	}

	amountPerTx := b.N / WORKERS

	server := NewFundServer(b.N)

	var wg sync.WaitGroup

	pizzaTime := false
	for i := 0; i < WORKERS; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for i := 0; i < amountPerTx; i++ {
				server.Transact(func(fund *model.Fund) {
					if fund.Balance() <= 10 {
						pizzaTime = true
						return
					}
					fund.Withdraw(1)
				})
			}
		}()
		if pizzaTime {
			break
		}
	}

	wg.Wait()

	balance := server.Balance()

	if balance != 10 {
		b.Error("Expected Balance to be 10, got", balance)
	}
}
