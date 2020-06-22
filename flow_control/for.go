package main

import "fmt"

func main() {
	sum := 0
	/*
		Go program language has only one looping construct, the for loop
		The for loop receives 3 components that is used to describe how the for loop will be executed.
		The first component is the initializer. This component is executed before the first iteration of the loop
		The second component is the condition. This component is verified after each iteration to check if the
	condition to stop the loop was reached
		The third component is the post statement. This component is executed after each iteration when the loop
	condition isn't reached.
	*/
	for i := 0; i < 10; i++ {
		sum += i
		fmt.Println(sum)
	}


	/*
		The initializer and post statement component are optional.
		If the components initializer and post statement aren't present, you have a while Go loop.
		The for loop can be used like a while loop on Go
	*/
	sum1 := 0
	for sum1 < 100 {
		sum1 += 10
		fmt.Println(sum1)
	}

	/*
		If you would like to execute a program forever, you can declare a for loop without condition.
		Every code inside the for loop will be executed forever
	*/
	sum2 := 0
	for {
		sum2 += 1
		fmt.Println(sum2)
	}


}