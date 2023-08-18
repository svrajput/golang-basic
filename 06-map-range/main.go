package main

import "fmt"

func main() {

	/* create a variable of map type
	has index and value of string type
	*/
	countryCapitalMap := make(map[string]string)

	/* insert key-value pairs in the map */
	countryCapitalMap["France"] = "Paris"
	countryCapitalMap["Italy"] = "Rome"
	countryCapitalMap["Japan"] = "Tokyo"
	countryCapitalMap["India"] = "New Delhi"

	/* Range is used to itrate over to print map values */
	for country := range countryCapitalMap {
		fmt.Println("Capital of", country, "is", countryCapitalMap[country])
	}

	/* test if entry is present in the map or not*/
	capital, ok := countryCapitalMap["United States"]

	/* if ok is true, entry is present otherwise entry is absent*/
	if ok {
		fmt.Println("Capital of United States is", capital)
	} else {
		fmt.Println("Capital of United States is not present")
	}

	fmt.Println("***********************************")
	fmt.Println("******** Delete Map Element ********")
	fmt.Println("***********************************")
	delete(countryCapitalMap, "France")

	/* print map */
	for country := range countryCapitalMap {
		fmt.Println("Capital of", country, "is", countryCapitalMap[country])
	}

}
