package main

import (  
    "fmt"
)

type Employee struct {  
    firstName string
    lastName  string
    age       int
    salary    int
}

func main() {

    //creating struct specifying field names
    emp1 := Employee{
        firstName: "Sunil",
        age:       40,
        salary:    5000,
        lastName:  "Rajput",
    }

    //creating struct without specifying field names
    emp2 := Employee{"John", "Do", 30, 800}

    fmt.Println("Employee 1", emp1)

    fmt.Println("Employee 2", emp2)
	
}