package main

import (
	"fmt"
	"sync"
)

var lock = new(sync.Mutex)
var instance *Help

type Help struct {
	name string
}

func (h Help) ShowName() string {
	return h.name
}

func GetInstance() *Help {
	if instance == nil {
		lock.Lock()
		defer lock.Unlock()
		instance = &Help{"Hello"}
	}
	return instance
}

func main() {
	x := GetInstance()
	fmt.Println(x.ShowName())
}
