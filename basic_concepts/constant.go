package main

import "fmt"

const Pi = 3.14

/*
Constants are declared like variables (local or global) using the operator var or :=, however with the const keyword.
A constant variable cannot be changed when the program is running. The value used to initialize a constant is fixed
A constant can be a boolean, string or any numerical variable
*/

func main() {
	const World = "世界"

	fmt.Println("Hello", World)
	fmt.Println("Happy", Pi, "Day")

	const Truth = true
	fmt.Println("Go rules?", Truth)
}
