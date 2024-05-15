package main

import "fmt"

func main() {
	//numbers := []int{1, 10, 15}
	//sum := sumup(numbers)
	// variadic function
	sum := sumup(1, 10, 15, 40, -5)
	numbers := []int{1, 10, 15}
	anotherSum := sumup(1, numbers...)

	fmt.Println(sum)
	fmt.Println(anotherSum)
}

// we want a slice
// func sumup(a []int) int {
// take any number of values, as long as they're int
// ... takes the arguments and turns it into a slice,
// so we don't need to create it ahead of time
// func sumup(a ...int) int {
// we take the first argument, and the rest get spread
func sumup(startingValue int, a ...int) int {
	sum := 0
	for _, v := range a {
		sum += v
	}
	return sum
}
