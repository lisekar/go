package main

import "fmt"

func main() {
	fmt.Println("Channels package main function.")
	messages := make(chan string)
	// go routine thread 2
	go func() {
		messages <- "ping"
	}()
	// main routine thread 1
	msg := <-messages
	fmt.Println(msg)
}
