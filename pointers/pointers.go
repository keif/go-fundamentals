package main

import "fmt"

func main() {
	age := 32 // regular variable, 32 stored in memory

	// declaring this is a pointer
	var agePointer *int
	// & "gets" the pointer
	agePointer = &age

	// * gets the value of the pointer
	fmt.Println("AgePointer:", agePointer)
	fmt.Println("AgePointer Value:", *agePointer)
	fmt.Println("Age:", age)

	// adultYears := getAdultYears(*age) // is equivalent, but we already created a pointer so we're using it
	//adultYears := getAdultYears(agePointer)
	//fmt.Println(adultYears)

	getAdultYears(agePointer)
	fmt.Println("Age:", age)
}

// this age is a copied value, in a different address in memory
// we're getting the address
// age is not being copied - but it's not really worth it, just an example on "how it works"
func getAdultYears(age *int) {
	// getting the value
	// if we want to update the pointer, we need to dereference the pointer
	// this is DIRECTLY mutating the value
	*age = *age - 18
}
