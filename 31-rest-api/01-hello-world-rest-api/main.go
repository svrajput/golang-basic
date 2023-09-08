package main

/*
1) Create First REST API greet  - Done
2) Handler Function and Request Multiplexer ( Router )  - Done
3) Code clean up create greet() and move response to new method.
4) Will see how to add response header
5) Response tag
*/

import (
	"fmt"
	"net/http"
)

func main() {

	http.HandleFunc("/greet", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello world rest response")
	})

	http.ListenAndServe("localhost:8080", nil)
}
