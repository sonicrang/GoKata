package main

import "fmt"

type IOperation interface {
	GetResult(a, b float64) float64
}

type AddOperation struct {
}

func (o AddOperation) GetResult(a, b float64) float64 {
	return a + b
}

type SubOperation struct {
}

func (o SubOperation) GetResult(a, b float64) float64 {
	return a - b
}

type MultOperation struct {
}

func (o MultOperation) GetResult(a, b float64) float64 {
	return a * b
}

type DivOperation struct {
}

func (o DivOperation) GetResult(a, b float64) float64 {
	if b == 0 {
		panic("can not divide 0")
	}
	return a / b
}

type Factory struct {
}

func (f Factory) GetFactory(oper string) IOperation {
	switch oper {
	case "+":
		return AddOperation{}
	case "-":
		return SubOperation{}
	case "*":
		return MultOperation{}
	case "/":
		return DivOperation{}
	default:
		return nil
	}
}
func main() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("Revocer: %v\n", err)
		}
	}()
	oper := Factory{}.GetFactory("/")
	fmt.Println(oper.GetResult(44, 0))
}
