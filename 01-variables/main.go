package main

import "fmt"

/* global variable declaration */
var a int = 20

func main() {
	// Static Type Declaration
	// Define variable and then assign value to variable
	var fname string
	fname = "Sunil "

	lname := "Rajput"

	// Static Type Declaration
	var age int = 30

	// print value of variable.
	fmt.Printf(" Name : %s %s \n", fname, lname)
	fmt.Println(" Age is : ", age)

	// Constants - variables.
	const LENGTH int = 10
	const WIDTH int = 5
	var area int

	area = LENGTH * WIDTH
	fmt.Println("value of area : ", area)

	/* local variable declaration in main function */
	var a, b, c int
	a = 10
	b = 20
	c = 0

	fmt.Printf("value of a in main() = %d\n", a)
	c = sum(a, b)
	fmt.Printf("value of c in main() = %d\n", c)

}

/* function to add two integers */
func sum(a, b int) int {
	fmt.Printf("value of a in sum() = %d\n", a)
	fmt.Printf("value of b in sum() = %d\n", b)

	return a + b
}

/* output of scope variables.
value of a in main() = 10
value of a in sum() = 10
value of b in sum() = 20
value of c in main() = 30
*/
