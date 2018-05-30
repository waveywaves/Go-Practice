package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func addtochannel(c chan int, i int) {
	defer wg.Done()
	c <- i
}

func main() {
	channel := make(chan int)
	defer close(channel)

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go addtochannel(channel, i)
		/*if the above is not a go routine we get the error below
		fatal error: all goroutines are asleep - deadlock!*/
	}

	for i := 0; i < 12; i++ {
		fmt.Println(<-channel)
	}
}
