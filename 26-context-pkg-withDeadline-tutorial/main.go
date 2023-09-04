package main

import (
	"context"
	"fmt"
	"time"
)

func main() {

	deadLine := time.Now().Add(2 * time.Millisecond)

	ctx, cancel := context.WithDeadline(context.Background(), deadLine)

	// Even though ctx will be expired, it is good practice to call its cancellation function in any case.
	// Failure to do so may keep the context and its parent alive longer than necessary.
	defer cancel()

	select {
	case <-ctx.Done():
		fmt.Println(ctx.Err())

	case <-time.After(1 * time.Second):
		fmt.Println(" over sleeped .. ")
	}

}
