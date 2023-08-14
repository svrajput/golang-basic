package main

import "fmt"

func main() {

	// Define variable and then assign value to variable
	var fname string
	fname = "Sunil "

	// Define varialble and assign value at the same time. Dynamic-typing capability
	lname := "Rajput"

	// print value of variable.
	fmt.Printf(" Full Name : %s %s \n", fname, lname)

	// Define Interger type of variable.
	var age int
	age = 30
	fmt.Println(" Age is : ", age)

	/*
		 you can also define more than one variable at the same time.
			var  c, ch byte;
			var  f, salary float32;
	*/

	// Constants - variables.
	const LENGTH int = 10
	const WIDTH int = 5
	var area int

	area = LENGTH * WIDTH
	fmt.Printf("value of area : %d", area)

}
