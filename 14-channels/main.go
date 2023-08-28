package main

import "fmt"

func main() {

	ch := make(chan int)

	go func(a, b int) {
		sum := a + b

		ch <- sum

	}(4, 2)

	r := <-ch
	// TODO: get the value computed from goroutine

	fmt.Printf("computed value %v\n", r)
}
