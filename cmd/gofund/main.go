package main

import (
	"fmt"

	"github.com/rxedu/go-fund"
)

func main() {
	fund := gofund.CreateFund()
	fmt.Printf("%v", fund.Balance())
}
