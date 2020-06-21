package main

import "fmt"

/*
A function can take zero or more arguments
In this example was created a function that receives 2 arguments and return a integer value
*/
func add(x int, y int) int {
	return x + y
}

func main() {
	value := add(42, 13)
	fmt.Printf("Example of sum using functions: ", value)
}
