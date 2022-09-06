package server

import (
	"fmt"

	"github.com/rxedu/go-fund/internal/model"
)

type FundServer struct {
	commands chan interface{}
	fund     *model.Fund
}

type WithdrawCommand struct {
	Amount int
}

type BalanceCommand struct {
	Response chan int
}

func NewFundServer(initialBalance int) *FundServer {
	server := &FundServer{
		commands: make(chan interface{}),
		fund:     model.NewFund(initialBalance),
	}
	go server.loop()
	return server
}

func (s FundServer) Balance() int {
	res := make(chan int)
	s.commands <- BalanceCommand{Response: res}
	return <-res
}

func (s FundServer) Withdraw(amount int) {
	s.commands <- WithdrawCommand{Amount: amount}
}

func (s FundServer) loop() {
	for command := range s.commands {
		switch cmd := command.(type) {
		case WithdrawCommand:
			s.fund.Withdraw(cmd.Amount)
		case BalanceCommand:
			balance := s.fund.Balance()
			cmd.Response <- balance
		default:
			panic(fmt.Sprintf("Unrecognized command: %v", command))
		}
	}
}
