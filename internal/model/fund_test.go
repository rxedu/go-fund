package model

import (
	"testing"
)

func BenchmarkFund(b *testing.B) {
	fund := NewFund(b.N)

	for i := 0; i < b.N; i++ {
		fund.Withdraw(1)
	}

	if fund.Balance() != 0 {
		b.Error("Expected Balance to be 0")
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
