package main

import "fmt"

type IFood interface {
	GetDesc() string
	GetPrice() float64
}

type Food struct {
}

func (o Food) GetDesc() string {
	return "菜单："
}

func (o Food) GetPrice() float64 {
	return 0
}

type Rice struct {
	IFood
	name  string
	price float64
}

func (o Rice) GetDesc() string {
	return fmt.Sprintf("%s %s", o.IFood.GetDesc(), o.name)
}

func (o Rice) GetPrice() float64 {
	return o.IFood.GetPrice() + o.price
}

func main() {
	f := Food{}
	r := Rice{f, "米饭", 20}
	fmt.Printf("%s\t价格：%f\n", r.GetDesc(), r.GetPrice())
	r1 := Rice{r, "炒饭", 30}
	fmt.Printf("%s\t价格：%f\n", r1.GetDesc(), r1.GetPrice())
}
