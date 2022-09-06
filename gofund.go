package gofund

import (
	"fmt"

	"github.com/rxedu/go-fund/internal/todo"
)

func PrintMessage() {
	fmt.Println(todo.GetMessage())
}
