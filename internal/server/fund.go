package server

import (
	"github.com/rxedu/go-fund/internal/model"
)

type FundServer struct {
	commands chan TransactionCommand
	fund     *model.Fund
}

type Transactor func(fund *model.Fund)

type TransactionCommand struct {
	Transactor Transactor
	Done       chan bool
}

func NewFundServer(initialBalance int) *FundServer {
	server := &FundServer{
		commands: make(chan TransactionCommand),
		fund:     model.NewFund(initialBalance),
	}
	go server.loop()
	return server
}

func (s FundServer) Transact(transactor Transactor) {
	command := TransactionCommand{
		Transactor: transactor,
		Done:       make(chan bool),
	}
	s.commands <- command
	<-command.Done
}

func (s FundServer) Balance() int {
	var balance int
	s.Transact(func(fund *model.Fund) {
		balance = fund.Balance()
	})
	return balance
}

func (s FundServer) Withdraw(amount int) {
	s.Transact(func(fund *model.Fund) {
		fund.Withdraw(amount)
	})
}

func (s FundServer) loop() {
	for command := range s.commands {
		command.Transactor(s.fund)
		command.Done <- true
	}
}
