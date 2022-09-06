package gofund

import "github.com/rxedu/go-fund/internal/model"

func CreateFund() *model.Fund {
	fund := model.NewFund(10)
	return fund
}
