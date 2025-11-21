package main

import (
	"fmt"
	"time"
)

func producer(ch chan int, count int) {
	for i := 0; i < count; i++ {
		ch <- i
		fmt.Printf("Producing %d \n", i)
		time.Sleep(time.Millisecond * 500)
	}
	close(ch)
}

func consumer(ch chan int) {
	for item := range ch { // receive until channel is closed
		fmt.Printf("Consuming %d \n", item)
		time.Sleep(time.Second * 1)
	}
}

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go producer(ch1, 10)
	consumer(ch1)

	go producer(ch2, 5)
	consumer(ch2)
}
