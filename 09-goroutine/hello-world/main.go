package main

import (
	"fmt"
	"time"
)

func printString(s string) {
   for i:= 0; i< 3; i++ {
	fmt.Println(s)
	time.Sleep(1 * time.Millisecond)
   }
 }

func main() {
	// Direct call
	printString("Direct call ")

	// TODO: write goroutine with different variants for function call.

	// goroutine function call
	go printString("called by goroutine-1 ")

	// goroutine with anonymous function
	go func() {
		printString("called by goroutine-2")
	}() 

	// goroutine with function value call
	fv := printString 
	go fv("called by goroutine-3")  	

	// wait for goroutines to end
	time.Sleep( 100 * time.Millisecond )

	fmt.Println("")
}
