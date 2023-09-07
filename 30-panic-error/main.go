package main

import (
	"fmt"
)

func panicTracer() {

	r := recover()

	if r != nil {
		fmt.Println(" Inside panic tracer ")
	}
}

func fullName(firstName *string, lastName *string) {
	defer panicTracer()

	if firstName == nil {
		panic("runtime error: first name cannot be nil")
	}

	if lastName == nil {
		panic("runtime error: last name cannot be nil")
	}

	fmt.Printf("%s %s\n", *firstName, *lastName)
}

func main() {
	firstName := "Sunil"
	// lastName := "Rajput"
	fullName(&firstName, nil)
}
