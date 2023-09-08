package main

/*
1) Create First REST API greet  - Done
2) Handler Function and Request Multiplexer ( Router )  - Done
3) Code clean up create greet() and move response to new method. - Done
4) lets create customers() end point  - Done
5) Will see how to add response header
6) Response tag for  -Done
*/

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type customer struct {
	Name    string `json:"full_name"`
	City    string `json:"city"`
	Zipcode string `json:"zip"`
}

func main() {

	// Defined route
	http.HandleFunc("/greet", greet)
	http.HandleFunc("/customers", customers)

	// starting server
	http.ListenAndServe("localhost:8080", nil)
}

func greet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello world rest response")
}

func customers(w http.ResponseWriter, r *http.Request) {
	customers := []customer{
		{"Sunil", "Dallas", "1111111"},
		{"Jon", "Dallas", "222222"},
	}

	w.Header().Add("Content-Type", "application/json")

	json.NewEncoder(w).Encode(customers)
}
