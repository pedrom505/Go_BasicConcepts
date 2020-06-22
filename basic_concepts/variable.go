package main

import "fmt"

//Global variable
var c, python, java = false, false, false

func changePython(){
	//Changing global variables
	python = true
}

func changeJava(){
	//Changing global variables
	java = true
}

func changeC(){
	//Changing global variables
	c = true
}


func main(){
	// Local Variable
	var i int // You can declare a variable mentioning your type
	/*
	Variables declared without an explicit initial value are given their zero value.
		The zero value is:
			0 for numeric types,
			false for the boolean type, and
			"" (the empty string) for strings.
	*/
	var j = 2 // If a initializer is present, the type can be omitted. The variable will take the type of the initializer
	k := 3 //Inside a function, the := short assignment statement can be used in place of a var declaration with implicit type.
	changePython()
	fmt.Println(i, j, c, k, python, java)
	changeJava()
	fmt.Println(i, j, c, k, python, java)
	//Changing local variables
	i=10
	changeC()
	fmt.Println(i, j, c, k, python, java)

}

