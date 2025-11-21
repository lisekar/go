package main

import (
	"fmt"
	"time"
)

func Greet(name string) {
	// Placeholder function
	fmt.Printf("Hello, %s!\n", name)
}
func main() {
	go Greet("World")
	time.Sleep(2 * time.Second)
	fmt.Println("Goroutine package main function.")
}
