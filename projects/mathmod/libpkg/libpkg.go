package libpkg

import "fmt"

func Greet(name string) {
	fmt.Printf("Hello, %s!\n", name)
	// return fmt.Sprintf("Greeting sent to %s", name)
}

func Add(a int, b int) int {
	return a + b
}
