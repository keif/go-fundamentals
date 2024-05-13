package main

import "fmt"

func main() {
	age := 32 // regular variable, 32 stored in memory

	fmt.Println("Age:", age)

	adultYears := getAdultYears(age)
	fmt.Println(adultYears)
}

// this age is a copied value, in a different address in memory
func getAdultYears(age int) int {
	return age - 18
}
