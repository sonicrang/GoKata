package main

import (
	"fmt"
	"time"
)

var channel = make(chan msg)

type msg struct {
	index   int
	content string
}

func Do(i int) {
	time.Sleep(1 * time.Second)
	channel <- msg{i, time.Now().String()}
}

func main() {
	go func() {
		for {
			select {
			case msg := <-channel:
				fmt.Printf("index:%d\tcontent:%s\n", msg.index, msg.content)
			default:
				// waiting
			}
		}
	}()
	fmt.Println("hi")
	for i := 0; i < 9; i++ {
		Do(i + 1)
	}
	fmt.Println("hello")
}
