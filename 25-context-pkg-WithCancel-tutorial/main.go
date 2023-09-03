package main

import (
	"context"
	"fmt"
	"log"
)

func main() {

	// Will use generator method to create goroutine which will increment counter varible
	// pass value of counter variable over channel
	// Once value of counter variable reaches to 5 we will call cancel function to terminate context.

	generator := func(ctx context.Context) <-chan int {
		dst := make(chan int)
		n := 1

		go func() {
			defer close(dst)
			for {
				select {
				case <-ctx.Done():
					log.Println(ctx.Err())
					return
				case dst <- n:
					n++
				}
			}
		}()

		return dst
	}

	//TODO : Create a cancellable context of type WithCancel
	ctx, cancel := context.WithCancel(context.Background())

	ch := generator(ctx)

	for n := range ch {
		fmt.Println(" value of counter n :", n)

		if n == 5 {
			cancel()
		}
	}

}
