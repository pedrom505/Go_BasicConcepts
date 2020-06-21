/*
package:
Usualy the package name is the same of the file where the code is writed, however
in this situation the chose the package name "main" because the program start running in the
package "main"
*/
package main

/*
Here are some examples showing how to import a package to be used in your code
You can import a package in your code using the command "import" following with the name of the package between "" or
import more than one package using the command "import" and mentioning the name of the packages that you need between ()
*/
import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println("My favorite number is", rand.Intn(10))
}
