package main

import (
	"fmt"
	"sync"
)

func main() {

	// TODO - modify code to print the value 1 as expected.
	// use waitgroup to do that.

	var cnt int
	var wg sync.WaitGroup

	wg.Add(1)

	go func() {
		defer wg.Done()
		cnt++
	}()

	wg.Wait()

	if cnt == 0 {
		fmt.Printf("value is %v \n", cnt)
	} else {
		fmt.Printf("value return is %v \n", cnt)
	}

	fmt.Println(" Done !! ")
}
