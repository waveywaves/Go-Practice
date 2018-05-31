package main

import "fmt"
import "sync"

var count int = 1

func say(s int, a chan int) {
	defer wg.Done()
	count++
	a <- s * count
}

var wg sync.WaitGroup

func main() {

	channel := make(chan int)

	//channel<-
	//<-channel
	//channel

	//var num int = 100

	wg.Add(2)
	go say(100, channel)
	go say(200, channel)
	//go say(num, channel)

	fmt.Println(<-channel)
}
