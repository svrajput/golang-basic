// package name
package main

//import dependency
import (
	"fmt"
	"math"
)

// constant is defined.
const strString string = "constant"

func main() {

	fmt.Println(strString)

	const n = 500000
	const d = 3e20 / n

	fmt.Println(int64(d))

	fmt.Println(math.Sin(n))

	// For loop example
	for j := 1; j <= 5; j++ {
		fmt.Println(j)
	}

	// If loop example
	if num := 9; num < 0 {
		fmt.Println(num, "is negative")
	} else if num < 10 {
		fmt.Println(num, "has 1 digit")
	} else {
		fmt.Println(num, "has multiple digits")
	}

	i := 10

	// switch statement example.
	switch i {

	case 1:
		fmt.Println(" one ")
	case 5:
		fmt.Println(" Five ")
	default:
		fmt.Println("It's ", i)
	}

}
