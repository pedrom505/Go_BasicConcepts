package main

import "fmt"

/*
	Go's return value may be named. When you are declaring the go's return for the function, you can use the return
field as the variable define field of the function
 */
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func main()  {
	fmt.Println(split(17))
}