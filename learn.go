package main

import "fmt"

type Person struct {
		Name string
		Age  int
	// 	func greeting() {
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
	greeting := func() {
		fmt.Println("Hello from a function!")
	}
	greeting()

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

	// range loop
	for index, value := range numSlice {
		fmt.Printf("Index: %d, Value: %d\n", index, value)
	}
	//while loop equivalent
	count := 1
	for count <= 5 {
		fmt.Println("Count:", count)
		count++
	}

	// infinite loop without break
	// for {
	// 	fmt.Println("This will run forever")
	// }

	// infinite loop with break
	for {
		fmt.Println("This will run once")
		break
	}


	// functions
	
	greet("Alice")
	
	result := add(5, 10)
	fmt.Println("Addition Result:", result)

	total := sum(15, 25)
	fmt.Println("Sum Result:", total)

	quotient, err := divide(10, 2)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Division Result:", quotient)
	}

	mult := multiply(2, 3, 4, 5, 8, 55, 90, 66)
	fmt.Println("Multiplication Result:", mult)

	closureCounter := makeCounter()
	fmt.Println("Closure Counter:", closureCounter())
	fmt.Println("Closure Counter:", closureCounter())
	fmt.Println("Closure Counter:", closureCounter())

	multiplyResult := func (a int) func(int) int {
		return func (b int) int {
			return a * b
		}	
	}(5)(6)

	fmt.Println("Multiply Result from nested anonymous function:", multiplyResult)

	sol := Multiplier(3)
	fmt.Println("Multiplier Result:", sol(10))

}

func greet(name string) {
	fmt.Printf("Hello, %s!\n", name)
}

func add(a, b int) int {
		return a + b
	}

func sum(a int, b int) int {
	return a + b
}

func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("division by zero")
	}
	return a / b, nil
}

// func with arbitrary number of arguments
func multiply(factors ...int) int {  // variadic function
	result := 1
	for _, factor := range factors {  // _ is used to ignore the index
		result *= factor
	}
	return result
}

// anonymous function example
var anonymousFunc = func(msg string) {
	fmt.Println(msg)
}

// closure example
func makeCounter() func() int {
	count := 0	
	return func() int {
		count++
		return count
	}
}

func Multiplier(factor int) func(int) int {
    return func(n int) int {
        return n * factor
    }
}



