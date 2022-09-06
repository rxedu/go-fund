package server

import (
	"fmt"

	"github.com/rxedu/go-fund/internal/model"
)

type FundServer struct {
	Commands chan interface{}
	fund     model.Fund
}

type WithdrawCommand struct {
	Amount int
}

type BalanceCommand struct {
	Response chan int
}

func NewFundServer(initialBalance int) *FundServer {
	server := &FundServer{
		Commands: make(chan interface{}),
		fund:     *model.NewFund(initialBalance),
	}
	go server.loop()
	return server
}

func (s FundServer) loop() {
	for command := range s.Commands {
		switch command.(type) {
		case WithdrawCommand:
			withdrawl := command.(WithdrawCommand)
			s.fund.Withdraw(withdrawl.Amount)
		case BalanceCommand:
			balanceReq := command.(BalanceCommand)
			balance := s.fund.Balance()
			balanceReq.Response <- balance
		default:
			panic(fmt.Sprintf("Unrecognized command: %v", command))
		}
	}
}
