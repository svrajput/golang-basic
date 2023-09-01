package main

import (
	"fmt"
	"sync"
	"time"
)

var sharedRsc = make(map[string]interface{})

func main() {

	var wg sync.WaitGroup

	m := sync.Mutex{}
	c := sync.NewCond(&m)

	wg.Add(1)
	go func() {
		defer wg.Done()

		c.L.Lock()
		for len(sharedRsc) == 0 {
			c.Wait()
		}
		fmt.Println(sharedRsc["rsc1"])
		c.L.Unlock()
	}()

	time.Sleep(2 * time.Second)
	fmt.Println("main goroutine ready")

	// TODO: writes changes to sharedRsc
	c.L.Lock()

	sharedRsc["rsc1"] = "foo"

	c.Signal()

	fmt.Println("main conditional variable goroutine signal send.")

	c.L.Unlock()

	wg.Wait()

}
