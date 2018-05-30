package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func say(s string) {
	defer wg.Done()
	for i := 0; i < 3; i++ {
		fmt.Println(s)
	}
}

func main() {
	wg.Add(1)
	go say("one")
	wg.Add(1)
	go say("two")
	wg.Add(1)
	go say("three")

	wg.Wait()
}
