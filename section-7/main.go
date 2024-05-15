package main

import "fmt"

func main() {
	//userNames := []string{} // creates empty array
	//userNames[0] = "Julie" // this will error, 0 doesn't exist
	//
	//userNames = append(userNames, "John") // new array
	//userNames = append(userNames, "Jane") // new array
	//// we could make go aware of the size of the array ahead of time to not re create the array
	//userNames := make([]string, 2) // creates an array of 2 empty elements
	userNames := make([]string, 2, 5) // creates an array of 2 empty elements, preallocated for more elements

	userNames = append(userNames, "John") // appended after the empty elements
	userNames = append(userNames, "Jane") // appended after the empty elements
	userNames[0] = "Julie"

	fmt.Println(userNames)
}
