package main

import (
	"fmt"
	"math"
)

type ICash interface {
	AcceptCash(float64) float64
}

type CashContext struct {
	ICash
}

func GetContext(s string) (cash CashContext, err error) {
	cash = CashContext{}
	switch s {
	case "normal":
		cash.ICash = NormalCash{}
	case "return":
		cash.ICash = ReturnCash{350, 100}
	case "rebate":
		cash.ICash = RebateCash{0.08}
	}
	return cash, err
}

type NormalCash struct {
}

func (c NormalCash) AcceptCash(value float64) float64 {
	return value
}

type ReturnCash struct {
	condition float64
	discount  float64
}

func (c ReturnCash) AcceptCash(value float64) float64 {
	return value - (math.Floor(value/c.condition) * c.discount)
}

type RebateCash struct {
	percent float64
}

func (c RebateCash) AcceptCash(value float64) float64 {
	return value * (1 - c.percent)
}

func main() {
	context, err := GetContext("normal")
	if err != nil {
		fmt.Println(err)
	} else {
		result := context.AcceptCash(800)
		fmt.Println(result)
	}

	context, err = GetContext("return")
	if err != nil {
		fmt.Println(err)
	} else {
		result := context.AcceptCash(800)
		fmt.Println(result)
	}

	context, err = GetContext("rebate")
	if err != nil {
		fmt.Println(err)
	} else {
		result := context.AcceptCash(800)
		fmt.Println(result)
	}
}
