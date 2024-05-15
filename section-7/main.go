package main

import "fmt"

type floatMap map[string]float64

func (f floatMap) output() {
	fmt.Println(f)
}

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

	// same problem as before, each addition creates a new map in memory
	//courseRatings := map[string]float64{}
	// same problem as above
	//courseRatings := make(map[string]float64)
	//courseRatings := make(map[string]float64, 3) // preallocate memory for three
	courseRatings := make(floatMap) // type alias
	courseRatings["go"] = 6.9
	courseRatings["java"] = 4.2
	courseRatings["python"] = 3.2 // still good, under the limit of the make function

	fmt.Println(courseRatings)

	for index, userName := range userNames {
		fmt.Println(index, userName, userNames[index])
	}

	for index, courseRating := range courseRatings {
		fmt.Println(index, courseRating, courseRatings[index])
	}
}
