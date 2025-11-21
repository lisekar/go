package main

import "fmt"

type Person struct {
	Name   string
	Age    int
	weight float32
}

func (p Person) Greeting() {
	fmt.Printf("Hello, my name is %s and I am %d years old.\n", p.Name, p.Age)
}

func (p *Person) SetWeight(w float32) {
	p.weight = w
}

func main() {
	p := Person{
		Name:   "Alice",
		Age:    30,
		weight: 65.5,
	}
	fmt.Println("Person Name:", p.Name)
	fmt.Println("Person Age:", p.Age)
	fmt.Println("Person Weight:", p.weight)
	p.Greeting()
	p.SetWeight(70.0)
	fmt.Println("Updated Person Weight:", p.weight)
}
