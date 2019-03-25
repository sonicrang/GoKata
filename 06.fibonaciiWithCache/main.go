package main

import "fmt"

var cacheMap = make(map[int]int, 10)

func Fibonacci(i int) int {
	value, ok := cacheMap[i]
	if ok {
		return value
	} else {
		switch {
		case i <= 0:
			return 0
		case i == 1:
			fallthrough
		case i == 2:
			cacheMap[i] = 1
			return cacheMap[i]
		default:
			cacheMap[i] = Fibonacci(i-1) + Fibonacci(i-2)
			return cacheMap[i]
		}
	}
}

func main() {
	// 1 1 2 3 5 8 13 21 34
	index := 1000
	fmt.Printf("index:%d\tvalue:%d", index, Fibonacci(index))
}
