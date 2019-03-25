package main

import "fmt"

func FibonacciNext() func() int {
	index := 0
	n1 := 0
	n2 := 0
	return func() int {
		index++
		switch index {
		case 1:
			n1 = 1
			return n1
		case 2:
			n2 = 1
			return n2
		default:
			n3 := n1 + n2
			n1 = n2
			n2 = n3
			return n1 + n2
		}
	}
}

func main() {
	// 1 1 2 3 5 8 13 21 34
	f := FibonacciNext()

	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}

}
