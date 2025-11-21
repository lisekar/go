package main

import "fmt"

type Greeter interface {
	Greet(name string)
}

type Person struct {
	Name string
}

func (p Person) Greet(name string) {
	fmt.Printf("Hello, %s! My name is %s.\n", name, p.Name)
}

func main() {
	var greeter Greeter
	greeter = Person{Name: "Alice"}
	greeter.Greet("Bob")
}
