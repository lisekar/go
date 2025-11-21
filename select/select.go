package main

import (
	"fmt"
	"time"
)

func sayHello(chName string) {
	// Placeholder function
	fmt.Printf("Hello from %s!\n", chName)
}

func main() {
	// fmt.Println("Select package main function.")
	ch1 := make(chan string)
	ch2 := make(chan string)
	go func() {
		time.Sleep(5 * time.Second)
		ch1 <- "Message from channel 1"
	}()
	go func() {
		// time.Sleep(5 * time.Second)
		ch2 <- "Message from channel 2"
	}()
	select {
	case msg1 := <-ch1:
		sayHello(msg1)
	case msg2 := <-ch2:
		sayHello(msg2)
	}
}
