package main

import "fmt"

//Global variable
var c, python, java bool

func changePython(){
	//Changing global variables
	python = true
}

func changeJava(){
	//Changing global variables
	java = true
}

func main(){
	// Local Variable
	var i int
	changePython()
	fmt.Println(i, c, python, java)
	changeJava()
	fmt.Println(i, c, python, java)
	//Changing local variables
	i=10
	fmt.Println(i, c, python, java)

}

