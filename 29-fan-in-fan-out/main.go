// Squaring numbers.

package main

import (
	"fmt"
	"sync"
)

func generator(nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		for _, n := range nums {
			out <- n
		}
		close(out)
	}()
	return out
}

func square(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- n * n
		}
		close(out)
	}()
	return out
}

func merge(cs ...<-chan int) <-chan int {
	out := make(chan int)
	var wg sync.WaitGroup

	output := func(c <-chan int) {
		for n := range c {
			out <- n
		}
		wg.Done()
	}

	wg.Add(len(cs))

	for _, n := range cs {
		go output(n)
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}

func main() {

	// stage 1 - Fan In
	in := generator(2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15)

	// stage 2 - Fan out
	ch1 := square(in)

	// stage 2   Fan out
	ch2 := square(in)

	// stage 3 - merge output of all channels. - Fan In
	output := merge(ch1, ch2)

	// stage 4 - print values
	for n := range output {
		fmt.Println("value received :", n)
	}

}
