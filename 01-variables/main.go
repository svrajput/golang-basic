package main

import "fmt"

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

	/*
		You can also define more than one variable at the same time.
		var  c, ch byte;
		var  f, salary float32;
	*/

	// Constants - variables.
	const LENGTH int = 10
	const WIDTH int = 5
	var area int

	area = LENGTH * WIDTH
	fmt.Println("value of area : ", area)

}
