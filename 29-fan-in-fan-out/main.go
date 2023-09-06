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

	// stage 1
	in := generator(2, 3)

	// stage 2
	ch1 := square(in)

	// stage 2
	ch2 := square(in)

	output := merge(ch1, ch2)

	for n := range output {
		fmt.Println("value received :", n)
	}

}
