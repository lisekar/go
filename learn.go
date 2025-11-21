package main

import "fmt"

type Person struct {
		Name string
		Age  int
	// 	func greet() {
	// 	fmt.Println("Hello!")
	// }
}


func main() {

	fmt.Println("Hello, Go!")

	// variable declaration
	var name string = "World"
	age := 25

	// var vote bool = true

	// zero values
	var id int
	var studentname string
	var ugOrpg bool

	fmt.Println("id:", id, "name:", studentname, "ugOrpg:", ugOrpg)

	fmt.Printf("Name: %s, Age: %d\n", name, age)

	// control structures
	if age >= 18 {
		fmt.Println("You are an adult.")
	} else {
		fmt.Println("You are a minor.")
	}

	// type conversion
	var score float64 = 95.5
	var intScore int = int(score)
	fmt.Printf("Score: %.2f, Int Score: %d\n", score, intScore)

	// constants
	const pi = 3.14
	fmt.Println("Value of pi:", pi)

	// pi := "3.14"
	// fmt.Println("Value of pi as string:", pi)


	//derived types	
	var numbers [5]int = [5]int{1, 2, 3, 4, 5}  // array - fixed size
	fmt.Println("Array:", numbers)

	// slice - dynamic size
	numSlice := []int{10, 20, 30, 40, 50}
	numSlice = append(numSlice, 60)
	fmt.Println("Slice:", numSlice)

	// map - key-value pairs
	studentScores := map[string]int{
		"Alice": 90,
		"Bob":   85,
	}
	studentScores["Charlie"] = 88
	fmt.Println("Alice's Score:", studentScores["Alice"])
	fmt.Println("Map:", studentScores)

	// struct - custom data type
	p := Person{Name: "David", Age: 30}
	fmt.Println("Struct:", p)
	// functions
	greet := func() {
		fmt.Println("Hello from a function!")
	}
	greet()

	// pointers
	var num int = 42
	var ptr *int = &num
	fmt.Println("Value:", num, "Pointer:", ptr, "Value via Pointer:", *ptr)
	*ptr = 100
	fmt.Println("Updated Value:", num)

	// if condition with short statement
	if n := 10; n%2 == 0 {
		fmt.Println(n, "is even")
	} else {
		fmt.Println(n, "is odd")
	}

	// inline variable declaration in
	if v := score; v > 90 {
		fmt.Println("High score!")
	} else if v >= 75 {
		fmt.Println("Good score!")
	} else {
		fmt.Println("Keep trying!")
	}

	// switch statement
	day := 3
	switch day {
	case 1:
		fmt.Println("Monday")
	case 2:
		fmt.Println("Tuesday")
	case 3:
		fmt.Println("Wednesday")
	default:
		fmt.Println("Another day")
	}

	// for loop
	for i := 1; i <= 5; i++ {
		fmt.Println("Iteration:", i)
	}
}