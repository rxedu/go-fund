package model

import (
	"sync"
	"testing"
)

const WORKERS = 10

func BenchmarkFund(b *testing.B) {
	if b.N < WORKERS {
		return
	}

	fund := NewFund(b.N)

	amountPerTx := b.N / WORKERS

	var wg sync.WaitGroup

	for i := 0; i < WORKERS; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for i := 0; i < amountPerTx; i++ {
				fund.Withdraw(1)
			}
		}()
	}

	wg.Wait()

	if fund.Balance() != 0 {
		b.Error("Expected Balance to be 0, got", fund.Balance())
	}
}

func TestFund(t *testing.T) {
	fund := NewFund(10)
	fund.Withdraw(4)

	got := fund.Balance()
	if got != 6 {
		t.Errorf("fund.Balance() = %d; want 6", got)
	}
}
