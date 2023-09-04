package main

import (
	"context"
	"fmt"
)

type userIDKey string
type database map[string]bool

var db database = database{
	"jane": true,
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	processRequest(ctx, "jane")
}

func processRequest(ctx context.Context, userid string) {
	// Created new context using parent context
	// send userID information to checkMemberShipStatus from database lookup.
	ctxv := context.WithValue(ctx, userIDKey("userIDKey"), userid)

	status := checkMemberShipStatus(ctxv)

	fmt.Printf("membership status of userid : %s : %v\n", userid, status)
}

// checkMemberShipStatus - takes context as input.
// extracts the user id information from context.
// spins a goroutine to do database lookup
// sends the result on the returned channel.
func checkMemberShipStatus(ctx context.Context) bool {
	userid := ctx.Value(userIDKey("userIDKey")).(string)
	status := db[userid]

	return status
}
