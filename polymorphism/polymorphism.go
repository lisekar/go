package main

type Animal interface {
	Speak() string
}

type Dog struct{}

func (d Dog) Speak() string {
	return "Woof!"
}

type Cat struct{}

func (c Cat) Speak() string {
	return "Meow!"
}

type Bird struct{}

func (b Bird) Speak() string {
	return "Chirp!"
}

func main() {
	var a Animal
	a = Dog{}
	println(a.Speak()) // Output: Woof!
	a = Cat{}
	println(a.Speak()) // Output: Meow!
	a = Bird{}
	println(a.Speak()) // Output: Chirp!

	animals := []Animal{Dog{}, Cat{}, Bird{}}
	for _, animal := range animals {
		println(animal.Speak())
	}

}
