package main

import "fmt"

func main() {
	var i, j int

	/*
	    **** Define an array of integer type ****
	   grade -- with elipsis &
	   NumbersArr as an array of integer type
	*/

	fmt.Println("***********************************")
	fmt.Println("**** Print Array elements *** ")
	fmt.Println("***********************************")
	grade := [...]int{2, 3, 5}
	var numbersArr = [5]int{1, 2, 3, 4, 5}

	/* output each array element's value */
	for i = 0; i < len(grade); i++ {
		fmt.Printf("grade[%d] = %d\n", i, grade[i])
	}

	/* output each array element's value */
	for j = 0; j < len(numbersArr); j++ {
		fmt.Printf("numbersArr[%d] = %d\n", j, numbersArr[j])
	}

	// ****  Define a variable of type slice ****
	fmt.Println("***********************************")
	fmt.Println("**** subsection of slice *** ")
	fmt.Println("***********************************")
	/* a slice of unspecified size */
	var numberSlice1 = make([]int, 3, 5)

	fmt.Printf("len=%d cap=%d slice=%v\n", len(numberSlice1),
		cap(numberSlice1), numberSlice1)

	/* create a slice has elements */
	var numberSlice2 = []int{0, 1, 2, 3, 4, 5, 6, 7, 8}

	/* add more than one element at a slice */
	numberSlice2 = append(numberSlice2, 9, 10, 11)

	for i = 0; i < len(numberSlice2); i++ {
		fmt.Printf("numberSlice2[%d] = %d\n", i, numberSlice2[i])
	}

	fmt.Println("***********************************")
	fmt.Println("**** subsection of slice *** ")
	fmt.Println("***********************************")
	fmt.Println(numberSlice2)

	number_1 := numberSlice2[2:5]
	fmt.Println(number_1)

	number_2 := numberSlice2[:2]
	fmt.Println(number_2)

	number_3 := numberSlice2[4:]
	fmt.Println(number_3)

}
