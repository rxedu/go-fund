package model

type Fund struct {
	balance int
}

func NewFund(initialBalance int) *Fund {
	return &Fund{balance: initialBalance}
}

func (fund *Fund) Balance() int {
	return fund.balance
}

func (fund *Fund) Withdraw(amount int) {
	fund.balance -= amount
}
