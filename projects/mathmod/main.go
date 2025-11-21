package main

import (
	"fmt"
	"mathmod/libpkg"

	"github.com/gorilla/mux"
	"github.com/lisekar/mathmodpublic"
)

func main() {
	router := mux.NewRouter() // This is a placeholder for the mathmod project main function.
	fmt.Println("Mathmod project initialized with router:", router)

	libpkg.Greet("Developer")
	sum := libpkg.Add(10, 20)
	fmt.Println("Sum from libpkg.Add:", sum)

	mult := mathmodpublic.Multiply(4, 5)
	fmt.Println("Multiplication from mathmodpublic.Multiply:", mult)

	if result, err := mathmodpublic.Divide(20, 4); err != nil {
		fmt.Println("Error in division:", err)
	} else {
		fmt.Println("Division from mathmodpublic.Divide:", result)
	}

	power := mathmodpublic.Power(2, 3)
	fmt.Println("Power from mathmodpublic.Power:", power)

}
